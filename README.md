This is a simple Go application that calls CoinMarketCap's API to get stats 
on any cryptocurrency you choose. It then uses your Twilio account to send
you an SMS with that information.

Please place your Twilio credentials in the **credentials.conf** file.

This program can be used with the *-coin* and *-info* flags.

If you do not specify flags, the defaults will be used. The default coin
returned is bitcoin while the default info level is price.

Example:

**./main -coin=litecoin -info=all** returns the following data:

```
Name: Litecoin 
Symbol: LTC 
PriceBtc: 0.0183727 
TotalSupply: 55201008.0 
PercentChange24H: -0.24 
ID: litecoin 
Rank: 6 
Two4HVolumeUsd: 503106000.0 
MarketCapUsd: 8778174719.0 
PriceUsd: 159.022 
LastUpdated: 1518546841 
AvailableSupply: 55201008.0 
MaxSupply: 84000000.0 
PercentChange1H: 0.77 
PercentChange7D: 18.81 
```

while **./main -coin=ripple -info=price** (or simply **./main -coin=ripple**)
returns the following:

```
The price of Ripple today is $1.02957
```
