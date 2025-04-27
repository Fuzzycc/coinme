# coinme
A simple tool to help you math how much your "Fantasy Kingdom Number 943" gold coins are worth in "Neighbor Kingdom of Fantasy Kingdom Number 943" copper coins. 

# The Problem
Have you ever drafted a fictional fantasy setting for a story, a TRPG campaign, etc., and found it a hassle to track how much a certain number of a certain coin is worth in another coin?

My character has 15 gold coins, each worth 10 silver coins, which in turn are worth 20 copper coins each. Some might say it's simple math, and my character has 3,000 copper coins... Whoa! That's an entire farming village's revenue in a year!

But the point is that my character is a capitalistic scroll-making wiz- ahem- The point is that, no matter how simple, you might not want to math it and need a quick answer.

# The Solution
Here's my copper, silver, and gold coins. I'll set their values at 1, 20, and 200, respectively.

Now, coinme 15 gold coins into copp- Oh, it's 3,000 copper coins, That's Great! (or G(o)lazingly Fast!) Now I can continue writing (plotting my campaign) without needing to interrupt my creative momentum to think if this amount of arbitrary money is plot-breaking or not.

# The Demonstration
Coinme revolves around Coins and Chains. The former are obvious, while the latter are a collection of coins, thus allowing you to keep multiple coin collections separate.
`coinme chain "Original Fantasy Continent"` This will create a new chain with the given name and return a unique ID for it.

`coinme coin 100 "Dwarvish Brewerdom Electrum Coin" 102` This will create a new coin with the given value, name, and an optional chain ID to which the coin would belong.

`coinme chain 102 1001 1002 1003 1004` This will add to the chain of ID 102 the coins of IDs 1001 to 1004.

`coinme remove chain 1102` This will remove the chain whose ID is 1102. The same for coin.

`coinme edit coin 1002 Name="Dwarvish Brewerdom Silvergold Coin" Value=110` This will update the coin with the ID 1002.

`coinme chain` This will return all created chains. Providing an ID will only return that chain, while a name will return all matches. The same for coin.

# The Future
Adding an interactive mode in the near future after initial setup.

Maybe wrapping the project into a website too, expanding it, then ambitiously calling it... a Fantasy Coinage Management System, or FCMS!!!!!!!!!!!
Jokes aside, a simple but user-friendly wrapper would be of great help to non-technical people, which is a major part of the envisioned user base.

# The Finale
Dot. Return. Return. Signature. Typewriter clickity clackity. Satisfying metallic DING.
