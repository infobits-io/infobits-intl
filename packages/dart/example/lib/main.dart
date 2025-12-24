import 'package:flutter/material.dart';
import 'package:infobits_intl/infobits_intl.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Infobits Intl Demo',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const HomePage(),
    );
  }
}

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  int _currentIndex = 0;

  static const List<Widget> _pages = [
    ContinentsTab(),
    CountriesTab(),
    CurrenciesTab(),
    LanguagesTab(),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _pages[_currentIndex],
      bottomNavigationBar: BottomNavigationBar(
        type: BottomNavigationBarType.fixed,
        currentIndex: _currentIndex,
        onTap: (index) => setState(() => _currentIndex = index),
        items: const [
          BottomNavigationBarItem(
            icon: Icon(Icons.public),
            label: 'Continents',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.flag),
            label: 'Countries',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.attach_money),
            label: 'Currencies',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.translate),
            label: 'Languages',
          ),
        ],
      ),
    );
  }
}

// ============================================================================
// Continents Tab
// ============================================================================

class ContinentsTab extends StatefulWidget {
  const ContinentsTab({super.key});

  @override
  State<ContinentsTab> createState() => _ContinentsTabState();
}

class _ContinentsTabState extends State<ContinentsTab> {
  final _searchController = TextEditingController();
  String _query = '';

  @override
  void dispose() {
    _searchController.dispose();
    super.dispose();
  }

  List<Continent> get _filtered {
    if (_query.isEmpty) return Continent.values;
    final lowerQuery = _query.toLowerCase();
    return Continent.values.where((c) {
      return c.name.toLowerCase().contains(lowerQuery) ||
          c.code.toLowerCase().contains(lowerQuery);
    }).toList();
  }

  @override
  Widget build(BuildContext context) {
    final items = _filtered;
    return Scaffold(
      appBar: AppBar(title: const Text('Continents'), centerTitle: true),
      body: Column(
        children: [
          Padding(
            padding: const EdgeInsets.all(16),
            child: TextField(
              controller: _searchController,
              onChanged: (v) => setState(() => _query = v),
              decoration: InputDecoration(
                hintText: 'Search continents...',
                prefixIcon: const Icon(Icons.search),
                border: OutlineInputBorder(borderRadius: BorderRadius.circular(12)),
                filled: true,
              ),
            ),
          ),
          Container(
            color: Theme.of(context).colorScheme.surfaceContainerHighest,
            padding: const EdgeInsets.symmetric(vertical: 12, horizontal: 16),
            child: const Row(
              children: [
                SizedBox(width: 60, child: Text('Code', style: TextStyle(fontWeight: FontWeight.bold))),
                Expanded(flex: 2, child: Text('Name', style: TextStyle(fontWeight: FontWeight.bold))),
                Expanded(child: Text('Countries', style: TextStyle(fontWeight: FontWeight.bold))),
                Expanded(child: Text('Languages', style: TextStyle(fontWeight: FontWeight.bold))),
              ],
            ),
          ),
          Expanded(
            child: ListView.builder(
              itemCount: items.length,
              itemBuilder: (context, index) {
                final continent = items[index];
                return Container(
                  padding: const EdgeInsets.symmetric(vertical: 12, horizontal: 16),
                  decoration: BoxDecoration(
                    border: Border(bottom: BorderSide(color: Colors.grey.shade300)),
                  ),
                  child: Row(
                    children: [
                      SizedBox(width: 60, child: Text(continent.code, style: const TextStyle(fontWeight: FontWeight.bold))),
                      Expanded(flex: 2, child: Text(continent.displayName(context))),
                      Expanded(child: Text('${continent.contries.length}')),
                      Expanded(child: Text('${continent.languages.toSet().length}')),
                    ],
                  ),
                );
              },
            ),
          ),
        ],
      ),
    );
  }
}

// ============================================================================
// Countries Tab
// ============================================================================

class CountriesTab extends StatefulWidget {
  const CountriesTab({super.key});

  @override
  State<CountriesTab> createState() => _CountriesTabState();
}

class _CountriesTabState extends State<CountriesTab> {
  final _searchController = TextEditingController();
  String _query = '';

  @override
  void dispose() {
    _searchController.dispose();
    super.dispose();
  }

  List<Country> get _filtered {
    if (_query.isEmpty) return Country.values;
    final lowerQuery = _query.toLowerCase();
    return Country.values.where((country) {
      return country.nativeName.toLowerCase().contains(lowerQuery) ||
          country.alpha2Code.toLowerCase().contains(lowerQuery) ||
          country.alpha3Code.toLowerCase().contains(lowerQuery) ||
          country.code.toLowerCase().contains(lowerQuery);
    }).toList();
  }

  @override
  Widget build(BuildContext context) {
    final items = _filtered;
    return Scaffold(
      appBar: AppBar(title: const Text('Countries'), centerTitle: true),
      body: Column(
        children: [
          Padding(
            padding: const EdgeInsets.all(16),
            child: TextField(
              controller: _searchController,
              onChanged: (v) => setState(() => _query = v),
              decoration: InputDecoration(
                hintText: 'Search countries...',
                prefixIcon: const Icon(Icons.search),
                border: OutlineInputBorder(borderRadius: BorderRadius.circular(12)),
                filled: true,
              ),
            ),
          ),
          // Header
          Container(
            color: Theme.of(context).colorScheme.surfaceContainerHighest,
            padding: const EdgeInsets.symmetric(vertical: 12, horizontal: 8),
            child: const Row(
              children: [
                SizedBox(width: 45, child: Text('Rect', style: TextStyle(fontWeight: FontWeight.bold))),
                SizedBox(width: 35, child: Text('Sq', style: TextStyle(fontWeight: FontWeight.bold))),
                SizedBox(width: 35, child: Text('Circ', style: TextStyle(fontWeight: FontWeight.bold))),
                SizedBox(width: 35, child: Text('', style: TextStyle(fontWeight: FontWeight.bold))),
                Expanded(flex: 2, child: Text('Name', style: TextStyle(fontWeight: FontWeight.bold))),
                Expanded(flex: 2, child: Text('Native Name', style: TextStyle(fontWeight: FontWeight.bold))),
                SizedBox(width: 50, child: Text('Alpha2', style: TextStyle(fontWeight: FontWeight.bold))),
                SizedBox(width: 55, child: Text('Alpha3', style: TextStyle(fontWeight: FontWeight.bold))),
                SizedBox(width: 55, child: Text('Num', style: TextStyle(fontWeight: FontWeight.bold))),
                SizedBox(width: 55, child: Text('Call', style: TextStyle(fontWeight: FontWeight.bold))),
                SizedBox(width: 40, child: Text('TLD', style: TextStyle(fontWeight: FontWeight.bold))),
                Expanded(flex: 2, child: Text('Continent', style: TextStyle(fontWeight: FontWeight.bold))),
                SizedBox(width: 45, child: Text('Curr', style: TextStyle(fontWeight: FontWeight.bold))),
              ],
            ),
          ),
          Expanded(
            child: ListView.builder(
              itemCount: items.length,
              itemBuilder: (context, index) {
                final country = items[index];
                return Container(
                  padding: const EdgeInsets.symmetric(vertical: 6, horizontal: 8),
                  decoration: BoxDecoration(
                    border: Border(bottom: BorderSide(color: Colors.grey.shade300)),
                  ),
                  child: Row(
                    children: [
                      SizedBox(width: 45, child: country.flag(height: 18)),
                      SizedBox(width: 35, child: Center(child: country.flag(size: 24, shape: FlagShape.square))),
                      SizedBox(width: 35, child: Center(child: country.flag(size: 24, shape: FlagShape.circle))),
                      SizedBox(width: 35, child: Text(country.emojiFlag, style: const TextStyle(fontSize: 18))),
                      Expanded(flex: 2, child: Text(country.displayName(context), overflow: TextOverflow.ellipsis)),
                      Expanded(flex: 2, child: Text(country.nativeName, overflow: TextOverflow.ellipsis)),
                      SizedBox(width: 50, child: Text(country.alpha2Code)),
                      SizedBox(width: 55, child: Text(country.alpha3Code)),
                      SizedBox(width: 55, child: Text(country.numericCodeString)),
                      SizedBox(width: 55, child: Text(country.callingCodeString)),
                      SizedBox(width: 40, child: Text(country.tld)),
                      Expanded(flex: 2, child: Text(country.continent.displayName(context), overflow: TextOverflow.ellipsis)),
                      SizedBox(width: 45, child: Text(country.currency.code)),
                    ],
                  ),
                );
              },
            ),
          ),
        ],
      ),
    );
  }
}

// ============================================================================
// Currencies Tab
// ============================================================================

class CurrenciesTab extends StatefulWidget {
  const CurrenciesTab({super.key});

  @override
  State<CurrenciesTab> createState() => _CurrenciesTabState();
}

class _CurrenciesTabState extends State<CurrenciesTab> {
  final _searchController = TextEditingController();
  String _query = '';

  @override
  void dispose() {
    _searchController.dispose();
    super.dispose();
  }

  List<Currency> get _filtered {
    if (_query.isEmpty) return Currency.values;
    final lowerQuery = _query.toLowerCase();
    return Currency.values.where((currency) {
      return currency.code.toLowerCase().contains(lowerQuery) ||
          currency.nativeName.toLowerCase().contains(lowerQuery) ||
          currency.symbol.toLowerCase().contains(lowerQuery);
    }).toList();
  }

  @override
  Widget build(BuildContext context) {
    final items = _filtered;
    return Scaffold(
      appBar: AppBar(title: const Text('Currencies'), centerTitle: true),
      body: Column(
        children: [
          Padding(
            padding: const EdgeInsets.all(16),
            child: TextField(
              controller: _searchController,
              onChanged: (v) => setState(() => _query = v),
              decoration: InputDecoration(
                hintText: 'Search currencies...',
                prefixIcon: const Icon(Icons.search),
                border: OutlineInputBorder(borderRadius: BorderRadius.circular(12)),
                filled: true,
              ),
            ),
          ),
          Container(
            color: Theme.of(context).colorScheme.surfaceContainerHighest,
            padding: const EdgeInsets.symmetric(vertical: 12, horizontal: 16),
            child: const Row(
              children: [
                SizedBox(width: 60, child: Text('Symbol', style: TextStyle(fontWeight: FontWeight.bold))),
                SizedBox(width: 60, child: Text('Code', style: TextStyle(fontWeight: FontWeight.bold))),
                Expanded(child: Text('Name', style: TextStyle(fontWeight: FontWeight.bold))),
                Expanded(child: Text('Plural', style: TextStyle(fontWeight: FontWeight.bold))),
              ],
            ),
          ),
          Expanded(
            child: ListView.builder(
              itemCount: items.length,
              itemBuilder: (context, index) {
                final currency = items[index];
                return Container(
                  padding: const EdgeInsets.symmetric(vertical: 12, horizontal: 16),
                  decoration: BoxDecoration(
                    border: Border(bottom: BorderSide(color: Colors.grey.shade300)),
                  ),
                  child: Row(
                    children: [
                      SizedBox(width: 60, child: Text(currency.symbol, style: const TextStyle(fontSize: 18, fontWeight: FontWeight.bold))),
                      SizedBox(width: 60, child: Text(currency.code)),
                      Expanded(child: Text(currency.displayName(context), overflow: TextOverflow.ellipsis)),
                      Expanded(child: Text(currency.nativeNamePlural, overflow: TextOverflow.ellipsis)),
                    ],
                  ),
                );
              },
            ),
          ),
        ],
      ),
    );
  }
}

// ============================================================================
// Languages Tab
// ============================================================================

class LanguagesTab extends StatefulWidget {
  const LanguagesTab({super.key});

  @override
  State<LanguagesTab> createState() => _LanguagesTabState();
}

class _LanguagesTabState extends State<LanguagesTab> {
  final _searchController = TextEditingController();
  String _query = '';

  @override
  void dispose() {
    _searchController.dispose();
    super.dispose();
  }

  List<Language> get _filtered {
    if (_query.isEmpty) return Language.values;
    final lowerQuery = _query.toLowerCase();
    return Language.values.where((language) {
      return language.code.toLowerCase().contains(lowerQuery) ||
          language.nativeName.toLowerCase().contains(lowerQuery);
    }).toList();
  }

  @override
  Widget build(BuildContext context) {
    final items = _filtered;
    return Scaffold(
      appBar: AppBar(title: const Text('Languages'), centerTitle: true),
      body: Column(
        children: [
          Padding(
            padding: const EdgeInsets.all(16),
            child: TextField(
              controller: _searchController,
              onChanged: (v) => setState(() => _query = v),
              decoration: InputDecoration(
                hintText: 'Search languages...',
                prefixIcon: const Icon(Icons.search),
                border: OutlineInputBorder(borderRadius: BorderRadius.circular(12)),
                filled: true,
              ),
            ),
          ),
          Container(
            color: Theme.of(context).colorScheme.surfaceContainerHighest,
            padding: const EdgeInsets.symmetric(vertical: 12, horizontal: 16),
            child: const Row(
              children: [
                SizedBox(width: 50, child: Text('Code', style: TextStyle(fontWeight: FontWeight.bold))),
                Expanded(child: Text('Name', style: TextStyle(fontWeight: FontWeight.bold))),
                Expanded(child: Text('Native Name', style: TextStyle(fontWeight: FontWeight.bold))),
                SizedBox(width: 70, child: Text('Dialects', style: TextStyle(fontWeight: FontWeight.bold))),
              ],
            ),
          ),
          Expanded(
            child: ListView.builder(
              itemCount: items.length,
              itemBuilder: (context, index) {
                final language = items[index];
                return Container(
                  padding: const EdgeInsets.symmetric(vertical: 12, horizontal: 16),
                  decoration: BoxDecoration(
                    border: Border(bottom: BorderSide(color: Colors.grey.shade300)),
                  ),
                  child: Row(
                    children: [
                      SizedBox(width: 50, child: Text(language.code.toUpperCase(), style: const TextStyle(fontWeight: FontWeight.bold))),
                      Expanded(child: Text(language.displayName(context), overflow: TextOverflow.ellipsis)),
                      Expanded(child: Text(language.nativeName, overflow: TextOverflow.ellipsis)),
                      SizedBox(width: 70, child: Text('${language.dialects.length}')),
                    ],
                  ),
                );
              },
            ),
          ),
        ],
      ),
    );
  }
}
