import 'dart:ui';

import 'package:flutter_test/flutter_test.dart';

import 'package:infobits_intl/infobits_intl.dart';
import 'package:infobits_intl/src/i18n/languages.g.dart';

void main() {
  test("Test that all translated language strings and keys follow ISO 639", () {
    LanguagesTranslationsDelegate.translations.forEach((key, translations) {
      translations.forEach((languageCode, languageName) {
        expect(languageName, isNotNull,
            reason: "Language $languageCode not found for $key");
        expect(languageName, isNotEmpty,
            reason: "Language $languageCode is empty for $key");
        expect(languageCode.length, 2,
            reason: "Language $languageCode is not of length 2 for $key");
      });
    });
  });

  test("Test that all translated language codes exsists in Languages enum", () {
    LanguagesTranslationsDelegate.translations.forEach((key, translations) {
      translations.forEach((languageCode, languageName) {
        final language = Language.fromCode(languageCode);
        expect(language, isNotNull,
            reason: "Language $languageCode not found for $key");
      });
    });
  });

  test("Test that all languages are translated", () {
    LanguagesTranslationsDelegate.translations.forEach((key, translations) {
      for (var language in Language.values) {
        final name = LanguagesTranslationsDelegate.getNameFromCodeAndLocale(
            language.code, Locale(key));

        expect(name, isNotNull,
            reason: "Language ${language.code} not found for translation $key");
        expect(name, isNotEmpty,
            reason: "Language ${language.code} is empty for translation $key");
      }
    });
  });
}
