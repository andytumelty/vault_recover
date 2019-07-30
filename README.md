Vault recover
=============

Extract master key from running unsealed vault.

Installation
============

`
pip install --user 'git+https://github.com/Vayu/vault_recover.git'
`

or alternatively

`
pip install --user 'git+ssh://git@github.com/Vayu/vault_recover.git'
`

Usage
=====

```
root@stretch:~# ./vault_recover.py --pid $(pidof vault) --keyring /root/vault-data/core/_keyring
INFO:root:Trying pid 3829
INFO:root:Scanning pid 3829 range 7fff58574000-7fff58595000 [stack];
INFO:root:Found 0 candidate pointers
INFO:root:Scanning pid 3829 range 7f5ab4042000-7f5ab42a2000 ;
INFO:root:Found 0 candidate pointers
INFO:root:Scanning pid 3829 range c420500000-c420600000 ;
INFO:root:Found 0 candidate pointers
INFO:root:Scanning pid 3829 range c420000000-c420500000 ;
INFO:root:Found 13 candidate pointers
INFO:root:Key c6932b0b230154b05c28bae8abf1c513a9ba8df61cd76c2133b39254c1d53510
INFO:root:Plaintext {"MasterKey":"xpMrCyMBVLBcKLroq/HFE6m6jfYc12whM7OSVMHVNRA=","Keys":[{"Term":1,"Version":1,"Value":"Smm6C2Pe3Wn3rYVHYQ5ILIJsH/Gs05Pv9M3yoHoTiBY=","InstallTime":"2019-07-29T22:17:46.159434884Z"}]}

andytumelty/vault-recover % go install shamir.go

andytumelty/vault-recover % shamir c6932b0b230154b05c28bae8abf1c513a9ba8df61cd76c2133b39254c1d53510
13a5955bf8ee1aeb01fa56ca81de107fac1d6409c99c5c6cf0637b5740ab0f81bb
4228ba05cf8b82201cee9e4599c74bd59bb5989986d377f6b8d2fa1c7ea6a5bbbe
deac06d255e92c03963f20d2e0049b7fde4288f4ce25d942fcb0b1102cdef64578
6c97d6c3e87f6bcd4ba4c0018143b9bb2544ab9fd72422689d7840ab680a8532c0
c16d1262cb607877040ba081bc5e01216e9cfd93451992dd52bb6d4b85dd54040d

root@stretch:~# vault seal
Vault is now sealed.

root@stretch:~# vault unseal 13a5955bf8ee1aeb01fa56ca81de107fac1d6409c99c5c6cf0637b5740ab0f81bb
Sealed: true
..
Unseal Progress: 1
root@stretch:~# vault unseal 4228ba05cf8b82201cee9e4599c74bd59bb5989986d377f6b8d2fa1c7ea6a5bbbe
Sealed: true
...
Unseal Progress: 2
root@stretch:~# vault unseal deac06d255e92c03963f20d2e0049b7fde4288f4ce25d942fcb0b1102cdef64578
Sealed: false
...
Unseal Progress: 0
```
