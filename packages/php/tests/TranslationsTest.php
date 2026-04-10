<?php

declare(strict_types=1);

namespace Infobits\Intl\Tests;

use Infobits\Intl\I18n\CapitalsTranslations;
use Infobits\Intl\I18n\ContinentsTranslations;
use Infobits\Intl\I18n\CountriesTranslations;
use Infobits\Intl\I18n\CurrenciesTranslations;
use Infobits\Intl\I18n\LanguagesTranslations;
use PHPUnit\Framework\TestCase;

final class TranslationsTest extends TestCase
{
    // ── Country translations ─────────────────────────────────────────

    public function testCountriesGet(): void
    {
        $name = CountriesTranslations::get('US', 'en');
        $this->assertNotNull($name);
        $this->assertNotEmpty($name);
    }

    public function testCountriesGetMissingCode(): void
    {
        $this->assertNull(CountriesTranslations::get('XX', 'en'));
    }

    public function testCountriesGetMissingLocale(): void
    {
        $this->assertNull(CountriesTranslations::get('US', 'zz_nonexistent'));
    }

    public function testCountriesAll(): void
    {
        $en = CountriesTranslations::all('en');
        $this->assertNotEmpty($en);
        $this->assertArrayHasKey('US', $en);
        $this->assertArrayHasKey('GB', $en);
    }

    public function testCountriesAllMissingLocale(): void
    {
        $this->assertEmpty(CountriesTranslations::all('zz_nonexistent'));
    }

    public function testCountriesLocales(): void
    {
        $locales = CountriesTranslations::locales();
        $this->assertNotEmpty($locales);
        $this->assertContains('en', $locales);
    }

    public function testCountriesMultipleLocales(): void
    {
        $locales = CountriesTranslations::locales();
        $this->assertGreaterThan(1, count($locales));
    }

    // ── Language translations ────────────────────────────────────────

    public function testLanguagesLocales(): void
    {
        $locales = LanguagesTranslations::locales();
        $this->assertNotEmpty($locales);
    }

    public function testLanguagesGet(): void
    {
        $locales = LanguagesTranslations::locales();
        if (!empty($locales)) {
            $locale = $locales[0];
            $all = LanguagesTranslations::all($locale);
            $this->assertNotEmpty($all);
        }
    }

    public function testLanguagesGetMissingLocale(): void
    {
        $this->assertNull(LanguagesTranslations::get('en', 'zz_nonexistent'));
    }

    // ── Currency translations ────────────────────────────────────────

    public function testCurrenciesGet(): void
    {
        $name = CurrenciesTranslations::get('USD', 'en');
        if ($name !== null) {
            $this->assertNotEmpty($name);
        }
    }

    public function testCurrenciesGetMissingCode(): void
    {
        $this->assertNull(CurrenciesTranslations::get('ZZZ', 'en'));
    }

    public function testCurrenciesLocales(): void
    {
        $locales = CurrenciesTranslations::locales();
        $this->assertNotEmpty($locales);
    }

    // ── Continent translations ───────────────────────────────────────

    public function testContinentsGet(): void
    {
        $name = ContinentsTranslations::get('EU', 'en');
        if ($name !== null) {
            $this->assertNotEmpty($name);
        }
    }

    public function testContinentsGetMissingCode(): void
    {
        $this->assertNull(ContinentsTranslations::get('XX', 'en'));
    }

    public function testContinentsLocales(): void
    {
        $locales = ContinentsTranslations::locales();
        $this->assertNotEmpty($locales);
    }

    // ── Capital translations ─────────────────────────────────────────

    public function testCapitalsLocales(): void
    {
        $locales = CapitalsTranslations::locales();
        $this->assertNotEmpty($locales);
    }

    public function testCapitalsGet(): void
    {
        $locales = CapitalsTranslations::locales();
        if (!empty($locales)) {
            $locale = $locales[0];
            $all = CapitalsTranslations::all($locale);
            $this->assertNotEmpty($all);
        }
    }

    public function testCapitalsGetMissingLocale(): void
    {
        $this->assertNull(CapitalsTranslations::get('US', 'zz_nonexistent'));
    }
}
