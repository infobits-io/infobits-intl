<?php

declare(strict_types=1);

namespace Infobits\Intl\Tests;

use Infobits\Intl\Country;
use Infobits\Intl\Currency;
use PHPUnit\Framework\TestCase;

final class CountryTest extends TestCase
{
    public function testCount(): void
    {
        $this->assertGreaterThan(200, count(Country::cases()));
    }

    public function testByAlpha2(): void
    {
        $us = Country::tryFrom('US');
        $this->assertNotNull($us);
        $this->assertSame('US', $us->alpha2Code());
        $this->assertSame('USA', $us->alpha3Code());
        $this->assertSame(840, $us->numericCode());
    }

    public function testKnownValues(): void
    {
        $us = Country::tryFrom('US');
        $this->assertNotNull($us);
        $this->assertSame('United States', $us->nativeName());
        $this->assertSame('Washington, D.C.', $us->capital());
        $this->assertSame('english', $us->mainLanguage());
        $this->assertContains('english', $us->languages());
        $this->assertSame('.us', $us->tld());
        $this->assertSame(1, $us->callingCode());
        $this->assertSame('northAmerica', $us->continent());
        $this->assertSame('usd', $us->currency());
    }

    public function testGermany(): void
    {
        $de = Country::tryFrom('DE');
        $this->assertNotNull($de);
        $this->assertSame('DEU', $de->alpha3Code());
        $this->assertSame(276, $de->numericCode());
        $this->assertSame('europe', $de->continent());
        $this->assertSame('eur', $de->currency());
    }

    public function testJapan(): void
    {
        $jp = Country::tryFrom('JP');
        $this->assertNotNull($jp);
        $this->assertSame('JPN', $jp->alpha3Code());
        $this->assertSame(392, $jp->numericCode());
        $this->assertSame('asia', $jp->continent());
    }

    public function testUnknownAlpha2ReturnsNull(): void
    {
        $this->assertNull(Country::tryFrom('XX'));
    }

    public function testAllHaveRequiredFields(): void
    {
        foreach (Country::cases() as $country) {
            $this->assertNotEmpty($country->alpha2Code(), "alpha2Code empty for {$country->name}");
            $this->assertNotEmpty($country->alpha3Code(), "alpha3Code empty for {$country->name}");
            $this->assertNotEmpty($country->nativeName(), "nativeName empty for {$country->name}");
            $this->assertNotEmpty($country->continent(), "continent empty for {$country->name}");
            $this->assertNotEmpty($country->currency(), "currency empty for {$country->name}");
        }
    }

    public function testFromAlpha3(): void
    {
        $us = Country::fromAlpha3('USA');
        $this->assertNotNull($us);
        $this->assertSame('US', $us->alpha2Code());

        $gb = Country::fromAlpha3('GBR');
        $this->assertNotNull($gb);
        $this->assertSame('GB', $gb->alpha2Code());
    }

    public function testFromAlpha3CaseInsensitive(): void
    {
        $us = Country::fromAlpha3('usa');
        $this->assertNotNull($us);
        $this->assertSame('US', $us->alpha2Code());
    }

    public function testFromAlpha3Unknown(): void
    {
        $this->assertNull(Country::fromAlpha3('XXX'));
    }

    public function testFromNumeric(): void
    {
        $us = Country::fromNumeric(840);
        $this->assertNotNull($us);
        $this->assertSame('US', $us->alpha2Code());

        $gb = Country::fromNumeric(826);
        $this->assertNotNull($gb);
        $this->assertSame('GB', $gb->alpha2Code());
    }

    public function testFromNumericUnknown(): void
    {
        $this->assertNull(Country::fromNumeric(999));
    }

    public function testLanguagesReturnsArray(): void
    {
        $us = Country::tryFrom('US');
        $this->assertNotNull($us);
        $this->assertIsArray($us->languages());
        $this->assertGreaterThan(0, count($us->languages()));
    }

    public function testEmojiFlagUS(): void
    {
        $us = Country::tryFrom('US');
        $this->assertNotNull($us);
        $this->assertSame("\u{1F1FA}\u{1F1F8}", $us->emojiFlag());
    }

    public function testEmojiFlagGB(): void
    {
        $gb = Country::tryFrom('GB');
        $this->assertNotNull($gb);
        $this->assertSame("\u{1F1EC}\u{1F1E7}", $gb->emojiFlag());
    }

    public function testEmojiFlagFR(): void
    {
        $fr = Country::tryFrom('FR');
        $this->assertNotNull($fr);
        $this->assertSame("\u{1F1EB}\u{1F1F7}", $fr->emojiFlag());
    }

    public function testEmojiFlagDE(): void
    {
        $de = Country::tryFrom('DE');
        $this->assertNotNull($de);
        $this->assertSame("\u{1F1E9}\u{1F1EA}", $de->emojiFlag());
    }

    public function testEmojiFlagJP(): void
    {
        $jp = Country::tryFrom('JP');
        $this->assertNotNull($jp);
        $this->assertSame("\u{1F1EF}\u{1F1F5}", $jp->emojiFlag());
    }

    public function testCurrencyReferencesExist(): void
    {
        $currencyCodes = array_map(fn (Currency $c) => strtolower($c->value), Currency::cases());

        $us = Country::tryFrom('US');
        $this->assertNotNull($us);
        $this->assertContains($us->currency(), $currencyCodes);

        $gb = Country::tryFrom('GB');
        $this->assertNotNull($gb);
        $this->assertContains($gb->currency(), $currencyCodes);
    }

    public function testContinentReferencesExist(): void
    {
        $continentIds = ['africa', 'antarctica', 'asia', 'europe', 'northAmerica', 'oceania', 'southAmerica'];
        foreach (Country::cases() as $country) {
            $this->assertContains(
                $country->continent(),
                $continentIds,
                "Country {$country->name} has invalid continent: {$country->continent()}"
            );
        }
    }
}
