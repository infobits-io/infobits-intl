import 'package:flutter_test/flutter_test.dart';

import 'package:infobits_intl/infobits_intl.dart';

void main() {
  test("Test that all currencies are valid", () {
    for (var currency in Currency.values) {
      expect(currency.code, isNotNull, reason: "Currency code is null");
      expect(currency.code.length, 3,
          reason: "Currency code is not of length 3");
      expect(currency.name, isNotNull, reason: "Currency name is null");
      expect(currency.name, isNotEmpty, reason: "Currency name is empty");
    }
  });

  test("Test that you can get a country from a ISO code", () {
    for (var currency in Currency.values) {
      var c = Currency.fromCode(currency.code);
      expect(c, isNotNull);
      expect(c?.code, currency.code);
    }
  });
}
