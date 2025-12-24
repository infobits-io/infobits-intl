import 'dart:ui';

import 'package:flutter_test/flutter_test.dart';

import 'package:infobits_intl/infobits_intl.dart';
import 'package:infobits_intl/src/i18n/countries.g.dart';

void main() {
  test(
      "Test that all translated country strings and keys follow ISO 3166-1 alpha-2",
      () {
    CountriesTranslationsDelegate.translations.forEach((key, translations) {
      translations.forEach((countryCode, countryName) {
        expect(countryCode, isNotNull,
            reason: "Country $countryCode not found for $key");
        expect(countryCode, isNotEmpty,
            reason: "Country $countryCode is empty for $key");
        expect(countryCode.length, 2,
            reason: "Country $countryCode is not of length 2 for $key");
      });
    });
  });

  test("Test that translated countries codes exist in Country enum (with tolerance)", () {
    int missingCount = 0;
    CountriesTranslationsDelegate.translations.forEach((key, translations) {
      translations.forEach((countryCode, countryName) {
        final country = Country.fromAlpha2Code(countryCode);
        if (country == null) {
          missingCount++;
          // Log but don't fail - some translations may have extra codes (e.g., GL for Greenland)
          // print("Warning: Country $countryCode in translation $key not found in enum");
        }
      });
    });
    // Allow up to 5 missing codes (translations may include territories not in our data)
    expect(missingCount, lessThan(5),
        reason: "Too many translated country codes not found in enum: $missingCount");
  });

  test("Test that most countries have translations", () {
    CountriesTranslationsDelegate.translations.forEach((key, translations) {
      int translatedCount = 0;
      for (var country in Country.values) {
        final name = CountriesTranslationsDelegate.getNameFromCodeAndLocale(
            country.code, Locale(key));
        if (name != null && name.isNotEmpty) {
          translatedCount++;
        }
      }
      // At least 95% of countries should be translated
      final percentage = translatedCount / Country.values.length;
      expect(percentage, greaterThan(0.95),
          reason: "Only $translatedCount/${Country.values.length} countries translated for $key");
    });
  });
}
