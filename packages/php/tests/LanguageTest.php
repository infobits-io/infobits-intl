<?php

declare(strict_types=1);

namespace Infobits\Intl\Tests;

use Infobits\Intl\Language;
use PHPUnit\Framework\TestCase;

final class LanguageTest extends TestCase
{
    public function testCount(): void
    {
        $this->assertGreaterThan(100, count(Language::cases()));
    }

    public function testByCode(): void
    {
        $en = Language::tryFrom('en');
        $this->assertNotNull($en);
        $this->assertSame('en', $en->code());
        $this->assertSame('english', $en->nativeName());
    }

    public function testFromCodeCaseInsensitive(): void
    {
        $en = Language::fromCode('EN');
        $this->assertNotNull($en);
        $this->assertSame('en', $en->code());
    }

    public function testFromCodeUnknown(): void
    {
        $this->assertNull(Language::fromCode('xx'));
    }

    public function testDefaultFlagCode(): void
    {
        $en = Language::tryFrom('en');
        $this->assertNotNull($en);
        $this->assertNotNull($en->defaultFlagCode());
    }

    public function testWithoutDefaultFlagCode(): void
    {
        $found = false;
        foreach (Language::cases() as $lang) {
            if ($lang->defaultFlagCode() === null) {
                $found = true;
                break;
            }
        }
        $this->assertTrue($found, 'Expected at least one language without a default flag code');
    }

    public function testKnownValues(): void
    {
        $fr = Language::tryFrom('fr');
        $this->assertNotNull($fr);
        $this->assertSame('fr', $fr->code());
        $this->assertSame('français', $fr->nativeName());

        $es = Language::tryFrom('es');
        $this->assertNotNull($es);
        $this->assertSame('es', $es->code());
        $this->assertSame('español', $es->nativeName());
    }
}
