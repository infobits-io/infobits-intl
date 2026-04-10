<?php

declare(strict_types=1);

namespace Infobits\Intl\Tests;

use Infobits\Intl\Flags;
use PHPUnit\Framework\TestCase;

final class FlagsTest extends TestCase
{
    public function testSvg(): void
    {
        $usFlag = Flags::svg('US');
        $this->assertNotNull($usFlag);
        $this->assertStringContainsString('<svg', $usFlag);
    }

    public function testSvgMultipleCountries(): void
    {
        $codes = ['US', 'GB', 'FR', 'DE', 'JP', 'BR', 'AU'];
        foreach ($codes as $code) {
            $flag = Flags::svg($code);
            $this->assertNotNull($flag, "Flag missing for {$code}");
            $this->assertStringContainsString('<svg', $flag, "Flag for {$code} is not valid SVG");
        }
    }

    public function testSvgCaseInsensitive(): void
    {
        $usFlag = Flags::svg('us');
        $this->assertNotNull($usFlag);
        $this->assertStringContainsString('<svg', $usFlag);
    }

    public function testSvgUnknown(): void
    {
        $this->assertNull(Flags::svg('XX'));
    }

    public function testAll(): void
    {
        $all = Flags::all();
        $this->assertGreaterThan(200, count($all));
        foreach ($all as $code => $svg) {
            $this->assertStringContainsString('<svg', $svg, "Flag for {$code} is not valid SVG");
        }
    }
}
