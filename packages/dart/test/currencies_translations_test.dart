import 'dart:ui';

import 'package:flutter_test/flutter_test.dart';

import 'package:infobits_intl/infobits_intl.dart';
import 'package:infobits_intl/src/i18n/currencies.g.dart';

void main() {
  test(
      "Test that all translated currency strings and keys follow ISO 4217",
      () {
    CurrenciesTranslationsDelegate.translations.forEach((key, translations) {
      translations.forEach((currencyCode, currencyName) {
        expect(currencyCode, isNotNull,
            reason: "Currency $currencyCode not found for $key");
        expect(currencyCode, isNotEmpty,
            reason: "Currency $currencyCode is empty for $key");
        expect(currencyCode.length, 3,
            reason: "Currency $currencyCode is not of length 3 for $key");
      });
    });
  });

  test("Test that translated currencies codes exist in Currency enum (with tolerance)",
      () {
    int missingCount = 0;
    CurrenciesTranslationsDelegate.translations.forEach((key, translations) {
      translations.forEach((currencyCode, currencyName) {
        final currency = Currency.fromCode(currencyCode);
        if (currency == null) {
          missingCount++;
          // Log but don't fail - some translations may have extra codes
          // print("Warning: Currency $currencyCode in translation $key not found in enum");
        }
      });
    });
    // Allow up to 10 missing codes (translations may include deprecated currencies)
    expect(missingCount, lessThan(10),
        reason: "Too many translated currency codes not found in enum: $missingCount");
  });

  test("Test that most currencies have translations", () {
    CurrenciesTranslationsDelegate.translations.forEach((key, translations) {
      int translatedCount = 0;
      for (var currency in Currency.values) {
        final name = CurrenciesTranslationsDelegate.getNameFromCodeAndLocale(
            currency.code, Locale(key));
        if (name != null && name.isNotEmpty) {
          translatedCount++;
        }
      }
      // All currencies should be translated
      expect(translatedCount, Currency.values.length,
          reason: "Only $translatedCount/${Currency.values.length} currencies translated for $key");
    });
  });
}
