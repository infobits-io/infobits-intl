<?php

// GENERATED CODE - DO NOT MODIFY BY HAND

declare(strict_types=1);

namespace Infobits\Intl;

enum Continent: string
{
    case Africa = 'AF';
    case Antarctica = 'AQ';
    case Asia = 'AS';
    case Europe = 'EU';
    case NorthAmerica = 'NA';
    case Oceania = 'OS';
    case SouthAmerica = 'SA';

    public function code(): string
    {
        return $this->value;
    }

    public function label(): string
    {
        return match ($this) {
            self::Africa => 'Africa',
            self::Antarctica => 'Antarctica',
            self::Asia => 'Asia',
            self::Europe => 'Europe',
            self::NorthAmerica => 'North America',
            self::Oceania => 'Oceania',
            self::SouthAmerica => 'South America',
        };
    }

    public static function fromCode(string $code): ?self
    {
        return self::tryFrom(strtoupper($code));
    }
}
