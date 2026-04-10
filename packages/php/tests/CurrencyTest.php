<?php

declare(strict_types=1);

namespace Infobits\Intl\Tests;

use Infobits\Intl\Currency;
use PHPUnit\Framework\TestCase;

final class CurrencyTest extends TestCase
{
    public function testCount(): void
    {
        $this->assertGreaterThan(100, count(Currency::cases()));
    }

    public function testByCode(): void
    {
        $usd = Currency::tryFrom('USD');
        $this->assertNotNull($usd);
        $this->assertSame('USD', $usd->code());
        $this->assertNotEmpty($usd->symbol());
        $this->assertSame('United States Dollar', $usd->nativeName());
        $this->assertSame('United States Dollars', $usd->nativeNamePlural());
    }

    public function testEuro(): void
    {
        $eur = Currency::tryFrom('EUR');
        $this->assertNotNull($eur);
        $this->assertSame('EUR', $eur->code());
        $this->assertSame('€', $eur->symbol());
    }

    public function testGBP(): void
    {
        $gbp = Currency::tryFrom('GBP');
        $this->assertNotNull($gbp);
        $this->assertSame('GBP', $gbp->code());
        $this->assertSame('£', $gbp->symbol());
    }

    public function testFromCodeCaseInsensitive(): void
    {
        $usd = Currency::fromCode('usd');
        $this->assertNotNull($usd);
        $this->assertSame('USD', $usd->code());
    }

    public function testFromCodeUnknown(): void
    {
        $this->assertNull(Currency::fromCode('ZZZ'));
    }

    public function testUnknownTryFromReturnsNull(): void
    {
        $this->assertNull(Currency::tryFrom('ZZZ'));
    }
}
