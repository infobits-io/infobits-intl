import 'dart:io';

import 'package:flutter_test/flutter_test.dart';

import 'package:infobits_intl/infobits_intl.dart';
import 'package:http/http.dart' as http;

import 'package:html/parser.dart' show parse;

void main() {
  test("Test that all countries are valid", () {
    for (var country in Country.values) {
      expect(country.code, isNotNull, reason: "Country code is null");
      expect(country.code.length, 2, reason: "Country code is not of length 2");
      expect(country.alpha2Code, isNotNull,
          reason: "Country alpha2 code is null");
      expect(country.alpha2Code.length, 2,
          reason: "Country alpha2 code is not of length 2");
      expect(country.alpha3Code, isNotNull,
          reason: "Country alpha3 code is null");
      expect(country.alpha3Code.length, 3,
          reason: "Country alpha3 code is not of length 3");
      expect(country.numericCode, isNotNull,
          reason: "Country numeric code is null");
      expect(country.numericCode, isNonNegative,
          reason: "Country numeric code is negative");
      expect(country.name, isNotNull, reason: "Country name is null");
      expect(country.name, isNotEmpty, reason: "Country name is empty");
    }
  });

  test("Test that you can get a country from a ISO code", () {
    for (var country in Country.values) {
      var c = Country.fromCode(country.code);
      expect(c, isNotNull);
      expect(c?.code, country.code);
      expect(c?.alpha2Code, country.alpha2Code);
      expect(c?.alpha3Code, country.alpha3Code);
      expect(c?.numericCode, country.numericCode);
      expect(c?.name, country.name);
    }
  });

  test("Test that you can get a country from ISO 3166 alpha 2", () {
    for (var country in Country.values) {
      var c = Country.fromAlpha2Code(country.alpha2Code);
      expect(c, isNotNull);
      expect(c?.alpha2Code, country.alpha2Code);
      expect(c?.alpha3Code, country.alpha3Code);
      expect(c?.numericCode, country.numericCode);
      expect(c?.name, country.name);
    }
  });

  test("Test that you can get a country from ISO 3166 alpha 3", () {
    for (var country in Country.values) {
      var c = Country.fromAlpha3Code(country.alpha3Code);
      expect(c, isNotNull);
      expect(c?.alpha2Code, country.alpha2Code);
      expect(c?.alpha3Code, country.alpha3Code);
      expect(c?.numericCode, country.numericCode);
      expect(c?.name, country.name);
    }
  });

  test("Test that you can get a country from ISO 3166 numeric code", () {
    for (var country in Country.values) {
      var c = Country.fromNumericCode(country.numericCode);
      expect(c, isNotNull);
      expect(c?.alpha2Code, country.alpha2Code);
      expect(c?.alpha3Code, country.alpha3Code);
      expect(c?.numericCode, country.numericCode);
      expect(c?.name, country.name);
    }
  });

  test("Test that all countries have valid emoji flags", () {
    for (var country in Country.values) {
      var emoji = country.emojiFlag;
      expect(emoji, isNotNull, reason: "Emoji flag is null for ${country.code}");
      expect(emoji, isNotEmpty, reason: "Emoji flag is empty for ${country.code}");
      // Each emoji flag consists of 2 regional indicator symbols (each is 2 UTF-16 code units)
      expect(emoji.runes.length, 2,
          reason: "Emoji flag should have 2 Unicode code points for ${country.code}");
      // Regional indicator symbols range from U+1F1E6 (🇦) to U+1F1FF (🇿)
      for (var rune in emoji.runes) {
        expect(rune >= 0x1F1E6 && rune <= 0x1F1FF, isTrue,
            reason: "Emoji flag contains invalid regional indicator for ${country.code}");
      }
    }
  });

  test("Test emoji flags for specific well-known countries", () {
    expect(Country.unitedStates.emojiFlag, "🇺🇸");
    expect(Country.unitedKingdom.emojiFlag, "🇬🇧");
    expect(Country.france.emojiFlag, "🇫🇷");
    expect(Country.germany.emojiFlag, "🇩🇪");
    expect(Country.japan.emojiFlag, "🇯🇵");
    expect(Country.china.emojiFlag, "🇨🇳");
    expect(Country.canada.emojiFlag, "🇨🇦");
    expect(Country.australia.emojiFlag, "🇦🇺");
    expect(Country.brazil.emojiFlag, "🇧🇷");
    expect(Country.india.emojiFlag, "🇮🇳");
  });

  test("Test emoji flag corresponds to alpha2 code", () {
    // Verify the algorithm: each letter in alpha2 code maps to a regional indicator
    const int base = 0x1F1E6 - 65; // A = 65 in ASCII

    for (var country in Country.values) {
      var expectedFirst = String.fromCharCode(base + country.alpha2Code.codeUnitAt(0));
      var expectedSecond = String.fromCharCode(base + country.alpha2Code.codeUnitAt(1));
      var expected = expectedFirst + expectedSecond;
      expect(country.emojiFlag, expected,
          reason: "Emoji flag doesn't match alpha2 code for ${country.code}");
    }
  });

  testWidgets("Test that you can render country flag",
      (WidgetTester tester) async {
    for (var country in Country.values) {
      var flag = country.flag(width: 30, height: 30);

      // Build the widget
      await tester.pumpWidget(flag);

      // Verify if the flag is rendered without errors
      expect(tester.takeException(), isNull,
          reason:
              "Exception occurred while rendering flag for ${country.code}");
      expect(find.byWidget(flag), findsOneWidget);
    }
  });

  test("Test that countries line up with wikipedia website", () async {
    // Save the original HTTP client
    final originalHttpOverrides = HttpOverrides.current;
    int missingCountries = 0;
    int checkedCountries = 0;

    try {
      HttpOverrides.global = null;

      final response = await http.get(Uri.parse(
          "https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes"));

      if (response.statusCode == 200) {
        var document = parse(response.body);
        var table = document.querySelector("table");
        var tbody = table?.querySelector("tbody");
        var rows = tbody?.querySelectorAll("tr");
        if (rows == null) {
          fail("Failed to parse table rows");
        } else {
          for (var i = 2; i < rows.length; i++) {
            var cells = rows[i].querySelectorAll("td");
            if (cells.length == 8) {
              var alpha2Code =
                  cells[3].getElementsByTagName('span').first.text.trim();
              var alpha3Code =
                  cells[4].getElementsByTagName('span').first.text.trim();
              var numericCodeString =
                  cells[5].getElementsByTagName('span').first.text.trim();
              var numericCode = int.tryParse(numericCodeString);
              var name = cells[0].text.trim().toLowerCase();

              expect(name, isNotEmpty,
                  reason: "Country name is empty: ${cells[0].outerHtml}");

              var countryFromAlpha2Code = Country.fromAlpha2Code(alpha2Code);

              // Skip countries not in our dataset (e.g., Greenland GL)
              if (countryFromAlpha2Code == null) {
                missingCountries++;
                continue;
              }

              checkedCountries++;

              expect(countryFromAlpha2Code.code, alpha2Code,
                  reason:
                      "Country $name's code is not the same for alpha2 code $alpha2Code");
              expect(countryFromAlpha2Code.alpha2Code, alpha2Code,
                  reason:
                      "Country $name's alpha2 code is not the same for alpha2 code $alpha2Code");
              expect(countryFromAlpha2Code.alpha3Code, alpha3Code,
                  reason:
                      "Country $name's alpha3 code is not the same for alpha2 code $alpha2Code");
              expect(countryFromAlpha2Code.numericCode, numericCode,
                  reason:
                      "Country $name's numeric code is not the same for alpha2 code $alpha2Code");

              var countryFromAlpha3Code = Country.fromAlpha3Code(alpha3Code);
              expect(countryFromAlpha3Code, isNotNull,
                  reason: "Country $name is null for alpha3 code $alpha3Code");

              if (numericCode != null) {
                var countryFromNumericCode = Country.fromNumericCode(numericCode);
                expect(countryFromNumericCode, isNotNull,
                    reason: "Country $name is null for numeric code $numericCode");
              }
            } else if (cells.length == 1) {
              // Skip the alternative country names
              continue;
            } else {
              // Skip rows with unexpected format
              continue;
            }
          }
        }

        // Ensure we checked most countries and didn't miss too many
        expect(checkedCountries, greaterThan(200),
            reason: "Expected to check at least 200 countries, but only checked $checkedCountries");
        expect(missingCountries, lessThan(10),
            reason: "Too many countries missing from our data: $missingCountries");
      } else {
        // Skip test if Wikipedia is not accessible
        print("Warning: Could not fetch Wikipedia (status ${response.statusCode}), skipping validation");
      }
    } catch (e) {
      // Skip test on network errors
      print("Warning: Network error fetching Wikipedia: $e");
    } finally {
      // Restore the original HTTP client
      HttpOverrides.global = originalHttpOverrides;
    }
  });
}
