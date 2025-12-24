// GENERATED CODE - DO NOT MODIFY BY HAND

import type { Continent } from './types';
import { countries, CountryCode } from './countries';

export enum ContinentCode {
  AF = "AF",
  AQ = "AQ",
  AS = "AS",
  EU = "EU",
  NA = "NA",
  OS = "OS",
  SA = "SA",
}

export const continents: Record<ContinentCode, Continent> = {
  [ContinentCode.AF]: {
    id: "africa",
    code: "AF",
    name: "Africa",
  },
  [ContinentCode.AQ]: {
    id: "antarctica",
    code: "AQ",
    name: "Antarctica",
  },
  [ContinentCode.AS]: {
    id: "asia",
    code: "AS",
    name: "Asia",
  },
  [ContinentCode.EU]: {
    id: "europe",
    code: "EU",
    name: "Europe",
  },
  [ContinentCode.NA]: {
    id: "northAmerica",
    code: "NA",
    name: "North America",
  },
  [ContinentCode.OS]: {
    id: "oceania",
    code: "OS",
    name: "Oceania",
  },
  [ContinentCode.SA]: {
    id: "southAmerica",
    code: "SA",
    name: "South America",
  },
} as const;

// Lookup functions
export function getContinentByCode(code: string): Continent | undefined {
  return continents[code.toUpperCase() as ContinentCode];
}

export function getCountriesInContinent(code: string): typeof countries[CountryCode][] {
  return Object.values(countries).filter(c => c.continent === code.toLowerCase());
}
