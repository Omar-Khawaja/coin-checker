This is a simple Go application that calls CoinMarketCap's API to get stats 
on any cryptocurrency you choose. It then uses your Twilio account to send
you an SMS with that information.

Please place your Twilio credentials in the **credentials.conf** file.

This program can be used with the *-coin* and *-info* flags.

If you do not use any flags, the price of bitcoin is returned by default.

Example:

**./main -coin=litecoin -info=all** returns the following data:

Name: Litecoin </br>
Symbol: LTC </br>
PriceBtc: 0.0183727 </br>
TotalSupply: 55201008.0 </br>
PercentChange24H: -0.24 </br>
ID: litecoin </br>
Rank: 6 </br>
Two4HVolumeUsd: 503106000.0 </br>
MarketCapUsd: 8778174719.0 </br>
PriceUsd: 159.022 </br>
LastUpdated: 1518546841 </br>
AvailableSupply: 55201008.0 </br>
MaxSupply: 84000000.0 </br>
PercentChange1H: 0.77 </br>
PercentChange7D: 18.81 </br>

while **./main -coin=ripple -info=price** (or simply **./main -coin=ripple**)
returns the following:

The price of Ripple today is $1.02957