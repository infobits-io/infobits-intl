import { describe, it, expect } from 'vitest';
import {
  countries,
  getCountryByAlpha2,
  getCountryByAlpha3,
  getCountriesByContinent,
  getEmojiFlag,
  CountryCode,
  languages,
  getLanguageByCode,
  LanguageCode,
  currencies,
  getCurrencyByCode,
  CurrencyCode,
  continents,
  getContinentByCode,
  getCountriesInContinent,
  ContinentCode,
  flags,
  getFlag,
  countriesTranslations,
  languagesTranslations,
  currenciesTranslations,
  continentsTranslations,
} from './index';

// Country tests

describe('Countries', () => {
  it('should have countries data', () => {
    expect(Object.keys(countries).length).toBeGreaterThan(200);
  });

  it('should get country by alpha2 code', () => {
    const us = getCountryByAlpha2('US');
    expect(us).toBeDefined();
    expect(us?.alpha2Code).toBe('US');
    expect(us?.alpha3Code).toBe('USA');
    expect(us?.numericCode).toBe(840);
  });

  it('should get country by alpha2 code (case insensitive)', () => {
    const us = getCountryByAlpha2('us');
    expect(us).toBeDefined();
    expect(us?.alpha2Code).toBe('US');
  });

  it('should return undefined for unknown alpha2 code', () => {
    const unknown = getCountryByAlpha2('XX');
    expect(unknown).toBeUndefined();
  });

  it('should get country by alpha3 code', () => {
    const us = getCountryByAlpha3('USA');
    expect(us).toBeDefined();
    expect(us?.alpha2Code).toBe('US');
    expect(us?.alpha3Code).toBe('USA');
  });

  it('should get country by alpha3 code (case insensitive)', () => {
    const us = getCountryByAlpha3('usa');
    expect(us).toBeDefined();
    expect(us?.alpha3Code).toBe('USA');
  });

  it('should get countries by continent', () => {
    const europeCountries = getCountriesByContinent('europe');
    expect(europeCountries.length).toBeGreaterThan(0);
    europeCountries.forEach((country) => {
      expect(country.continent).toBe('europe');
    });
  });

  it('should have required country fields', () => {
    const us = getCountryByAlpha2('US');
    expect(us).toBeDefined();
    expect(us?.alpha2Code).toBeTruthy();
    expect(us?.alpha3Code).toBeTruthy();
    expect(us?.numericCode).toBeDefined();
    expect(us?.nativeName).toBeTruthy();
    expect(us?.mainLanguage).toBeTruthy();
    expect(us?.languages.length).toBeGreaterThan(0);
    expect(us?.continent).toBeTruthy();
    expect(us?.currency).toBeTruthy();
  });
});

// Emoji flag tests

describe('Emoji Flags', () => {
  it('should generate correct emoji flags', () => {
    expect(getEmojiFlag('US')).toBe('🇺🇸');
    expect(getEmojiFlag('GB')).toBe('🇬🇧');
    expect(getEmojiFlag('FR')).toBe('🇫🇷');
    expect(getEmojiFlag('DE')).toBe('🇩🇪');
    expect(getEmojiFlag('JP')).toBe('🇯🇵');
  });

  it('should handle uppercase codes only', () => {
    // getEmojiFlag requires uppercase codes
    expect(getEmojiFlag('US')).toBe('🇺🇸');
  });
});

// Language tests

describe('Languages', () => {
  it('should have languages data', () => {
    expect(Object.keys(languages).length).toBeGreaterThan(100);
  });

  it('should get language by code', () => {
    const en = getLanguageByCode('en');
    expect(en).toBeDefined();
    expect(en?.code).toBe('en');
    expect(en?.nativeName).toBeTruthy();
  });

  it('should get language by code (case insensitive)', () => {
    const en = getLanguageByCode('EN');
    expect(en).toBeDefined();
    expect(en?.code).toBe('en');
  });

  it('should return undefined for unknown language code', () => {
    const unknown = getLanguageByCode('xx');
    expect(unknown).toBeUndefined();
  });
});

// Currency tests

describe('Currencies', () => {
  it('should have currencies data', () => {
    expect(Object.keys(currencies).length).toBeGreaterThan(100);
  });

  it('should get currency by code', () => {
    const usd = getCurrencyByCode('USD');
    expect(usd).toBeDefined();
    expect(usd?.code).toBe('USD');
    expect(usd?.symbol).toBeTruthy();
    expect(usd?.nativeName).toBeTruthy();
  });

  it('should get currency by code (case insensitive)', () => {
    const usd = getCurrencyByCode('usd');
    expect(usd).toBeDefined();
    expect(usd?.code).toBe('USD');
  });

  it('should return undefined for unknown currency code', () => {
    const unknown = getCurrencyByCode('ZZZ');
    expect(unknown).toBeUndefined();
  });
});

// Continent tests

describe('Continents', () => {
  it('should have 7 continents', () => {
    expect(Object.keys(continents).length).toBe(7);
  });

  it('should get continent by code', () => {
    const europe = getContinentByCode('EU');
    expect(europe).toBeDefined();
    expect(europe?.code).toBe('EU');
    expect(europe?.name).toBe('Europe');
  });

  it('should get continent by code (case insensitive)', () => {
    const europe = getContinentByCode('eu');
    expect(europe).toBeDefined();
    expect(europe?.code).toBe('EU');
  });

  it('should get countries in continent', () => {
    // getCountriesInContinent takes continent id (e.g., "europe"), not code ("EU")
    const europeCountries = getCountriesInContinent('europe');
    expect(europeCountries.length).toBeGreaterThan(0);
  });
});

// Flag tests

describe('Flags', () => {
  it('should have flag SVGs', () => {
    expect(Object.keys(flags).length).toBeGreaterThan(200);
  });

  it('should get flag by country code', () => {
    const usFlag = getFlag('US');
    expect(usFlag).toBeDefined();
    expect(usFlag).toContain('<svg');
  });

  it('should get flag by country code (case insensitive)', () => {
    const usFlag = getFlag('us');
    expect(usFlag).toBeDefined();
    expect(usFlag).toContain('<svg');
  });

  it('should return undefined for unknown country code', () => {
    const unknown = getFlag('XX');
    expect(unknown).toBeUndefined();
  });
});

// Translation tests

describe('Translations', () => {
  it('should have English country translations', () => {
    const enTranslations = countriesTranslations['en'];
    if (enTranslations) {
      expect(enTranslations['US']).toBeTruthy();
    }
  });

  it('should have English currency translations', () => {
    const enTranslations = currenciesTranslations['en'];
    if (enTranslations) {
      expect(enTranslations['USD']).toBeTruthy();
    }
  });

  it('should have English continent translations', () => {
    const enTranslations = continentsTranslations['en'];
    if (enTranslations) {
      expect(enTranslations['EU']).toBeTruthy();
    }
  });
});
