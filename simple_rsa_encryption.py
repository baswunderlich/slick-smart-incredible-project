import rsa

if __name__ == '__main__':

  (pubkey, privkey) = rsa.newkeys(512, poolsize=4)

  message = b'Say hi!'
  crypto = rsa.encrypt(message, pubkey)

  PRIV_KEY_DST = 'your path to file'
  with open(PRIV_KEY_DST, 'wb+') as f:
    pk = rsa.PrivateKey.save_pkcs1(privkey, format='PEM')
    f.write(pk)

  print(crypto)
  print(rsa.decrypt(crypto=crypto, priv_key=privkey))
  print(rsa.sign(priv_key=privkey, message=message, hash_method='SHA-1'))