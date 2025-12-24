import 'dart:ui';

import 'package:flutter_test/flutter_test.dart';

import 'package:infobits_intl/infobits_intl.dart';
import 'package:infobits_intl/src/i18n/continents.g.dart';

void main() {
  test(
      "Test that all translated continent strings and keys follow ISO 3166-1 alpha-2",
      () {
    ContinentsTranslationsDelegate.translations.forEach((key, translations) {
      translations.forEach((continentCode, continentName) {
        expect(continentCode, isNotNull,
            reason: "Continent $continentCode not found for $key");
        expect(continentCode, isNotEmpty,
            reason: "Continent $continentCode is empty for $key");
        expect(continentCode.length, 2,
            reason: "Continent $continentCode is not of length 2 for $key");
      });
    });
  });

  test("Test that all translated continent codes exsists in Continent enum",
      () {
    ContinentsTranslationsDelegate.translations.forEach((key, translations) {
      translations.forEach((continentCode, continentName) {
        final continent = Continent.fromCode(continentCode);
        expect(continent, isNotNull,
            reason: "Continent $continentCode not found for $key");
      });
    });
  });

  test("Test that all continents are translated", () {
    ContinentsTranslationsDelegate.translations.forEach((key, translations) {
      for (var continent in Continent.values) {
        final name = ContinentsTranslationsDelegate.getNameFromCodeAndLocale(
            continent.code, Locale(key));

        expect(name, isNotNull,
            reason:
                "Continent ${continent.code} not found for translation $key");
        expect(name, isNotEmpty,
            reason:
                "Continent ${continent.code} is empty for translation $key");
      }
    });
  });
}
