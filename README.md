# api_latencies


This is a test project to obtain the max USD value that can be processed in 50ms, 60ms, 90ms, 1000ms?

Data provided were:
1. api_latencies.json containing the time to process a transaction for a specific country using the bankCountryCode
2. transactions.csv containing transaction info like id, amount and bankCountryCode

From my findings and results obtained, It took the following xxxUSD for xxxms:

4139.43USD for 50ms (5 transactions)

4654.24USD for 60ms	 (6 transactions)

5763.62USD for 90ms	 (9 transactions)

10403.15USD for 1000ms (99 transactions) 


From the question, I observed it is a linear problem. I sorted the transactions using two(2) conditional filters i.e time latency and amount.
Though, there is still something fishing that I have not figured out.

Disclaimer: I am not absolutely confident of this result.


Thanks.
