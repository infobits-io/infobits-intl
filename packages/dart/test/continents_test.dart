import 'package:flutter_test/flutter_test.dart';

import 'package:infobits_intl/infobits_intl.dart';

void main() {
  test("Test that all continents are valid", () {
    for (var continent in Continent.values) {
      expect(continent.code, isNotNull, reason: "Continent code is null");
      expect(continent.code.length, 2,
          reason: "Continent code is not of length 2");
      expect(continent.name, isNotNull, reason: "Continent name is null");
      expect(continent.name, isNotEmpty, reason: "Continent name is empty");
    }
  });

  test("Test that you can get a continent from a ISO code", () {
    for (var continent in Continent.values) {
      var c = Continent.fromCode(continent.code);
      expect(c, isNotNull);
      expect(c?.code, continent.code);
    }
  });

  test("Test that you can get a continent from a english name", () {
    for (var continent in Continent.values) {
      var c = Continent.fromName(continent.name);
      expect(c, isNotNull);
      expect(c?.name, continent.name);
    }
  });
}
