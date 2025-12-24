import 'package:flutter_test/flutter_test.dart';

import 'package:infobits_intl/infobits_intl.dart';

void main() {
  test("Test that all languages are valid", () {
    for (var language in Language.values) {
      expect(language.code, isNotNull, reason: "Language code is null");
      expect(language.code.length, 2,
          reason: "Language code is not of length 2");
      expect(language.name, isNotNull, reason: "Language name is null");
      expect(language.name, isNotEmpty, reason: "Language name is empty");
    }
  });

  test("Test that you can get a language from a ISO 639 code", () {
    for (var language in Language.values) {
      var l = Language.fromCode(language.code);
      expect(l, isNotNull);
      expect(l?.code, language.code);
    }
  });
}
