<?php

declare(strict_types=1);

namespace Infobits\Intl\Tests;

use Infobits\Intl\Continent;
use PHPUnit\Framework\TestCase;

final class ContinentTest extends TestCase
{
    public function testCount(): void
    {
        $this->assertSame(7, count(Continent::cases()));
    }

    public function testByCode(): void
    {
        $europe = Continent::tryFrom('EU');
        $this->assertNotNull($europe);
        $this->assertSame('EU', $europe->code());
        $this->assertSame('Europe', $europe->label());
    }

    public function testAllLabels(): void
    {
        $expected = [
            'AF' => 'Africa',
            'AQ' => 'Antarctica',
            'AS' => 'Asia',
            'EU' => 'Europe',
            'NA' => 'North America',
            'OS' => 'Oceania',
            'SA' => 'South America',
        ];

        foreach ($expected as $code => $label) {
            $continent = Continent::tryFrom($code);
            $this->assertNotNull($continent, "Continent {$code} not found");
            $this->assertSame($label, $continent->label(), "Label mismatch for {$code}");
        }
    }

    public function testFromCodeCaseInsensitive(): void
    {
        $europe = Continent::fromCode('eu');
        $this->assertNotNull($europe);
        $this->assertSame('EU', $europe->code());
    }

    public function testFromCodeUnknown(): void
    {
        $this->assertNull(Continent::fromCode('XX'));
    }
}
