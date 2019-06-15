
# Nanobitcoin-Go

A golang implementation of [Nanobitcoin](https://github.com/acolytec3/nanobitcoin)

## Features

* A dumb http server that gives you a fake blockchain

## To-do list

* Write a skeleton for tests
* Port the block schema from Nanobitcoin
* Create methods to add transactions to a block
* Port brute-force transaction validation from Nanobitcoin
* Define PoA consensus algorithm - to replace proof of work cuz even Go isn't fast enough for that nonsense
* Make the note actually produce blocks on some regular basis
* Integrate a database (maybe [bbolt](https://github.com/etcd-io/bbolt)) to store the blocks
* Add routes that match [Nanobitcoin-Blockchain-Explorer](https://github.com/acolytec3/blockchain-explorer)

<table border="0"><tr>  <td><a href="https://gittron.me/bots/0xddab05da7f11acf322db7ac1c4f671ea"><img src="https://s3.amazonaws.com/od-flat-svg/0xddab05da7f11acf322db7ac1c4f671ea.png" alt="gittron" width="50"/></a></td><td><a href="https://gittron.me/bots/0xddab05da7f11acf322db7ac1c4f671ea">Buidl Nanobitcoin-Go with a Support Bot from Gittron</a></td></tr></table>
