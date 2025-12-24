// GENERATED CODE - DO NOT MODIFY BY HAND

export interface Country {
  readonly id: string;
  readonly alpha2Code: string;
  readonly alpha3Code: string;
  readonly numericCode: number;
  readonly nativeName: string;
  readonly capital: string;
  readonly mainLanguage: string;
  readonly languages: readonly string[];
  readonly tld: string;
  readonly callingCode: number;
  readonly continent: string;
  readonly currency: string;
}

export interface LanguageDialect {
  readonly id: string;
  readonly code: string;
  readonly nativeName: string;
  readonly flagCode?: string;
}

export interface Language {
  readonly id: string;
  readonly code: string;
  readonly nativeName: string;
  readonly dialects: readonly LanguageDialect[];
  readonly defaultFlagCode?: string;
}

export interface Currency {
  readonly id: string;
  readonly code: string;
  readonly nativeName: string;
  readonly nativeNamePlural: string;
  readonly symbol: string;
}

export interface Continent {
  readonly id: string;
  readonly code: string;
  readonly name: string;
}
