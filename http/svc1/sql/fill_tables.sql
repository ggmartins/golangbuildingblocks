-- bcrypt rounds 15
-- thanks to:
-- https://www.4devs.com.br/gerador_de_cpf
-- https://bcrypt-generator.com/

-- secret QaeZC minibank1
-- secret psTVK minibank2
-- secret FYoNO minibank3
INSERT INTO accounts (id, name, cpf, secret, balance, created_at) VALUES (1001, 'Steve Ray Vaughan', '40956406009', '$2a$15$aHPMf0.R/AhT1KWaE9AVNO0OuEY6av/NLYFHiObWN3Fs5oRuQaeZC', '800000.0000', '2020-10-10 19:17:12-05');
INSERT INTO accounts (id, name, cpf, secret, balance, created_at) VALUES (1002, 'Samantha Fish', '64857547090', '$2a$15$qlHpA/QvOLwcrJ1m0NA6eu8phqJ6zWinW1wArF9eph9K2uZMpsTVK', '123.1230', '2022-02-28 12:00:50-08');
INSERT INTO accounts (id, name, cpf, secret, balance, created_at) VALUES (1003, 'Joe Satriani', '66984278023', '$2a$15$tJHEa9RtE20dIVCzGod5seZ4wViJxzy49a4Ftepc.nPEx0A1FYoNO', '0.0', '2022-01-01 00:06:00-04');
INSERT INTO accounts (id, name, cpf, secret, balance, created_at) VALUES (1004, 'Jimmy Hendrix', '90450195074', '$2a$15$aHPMf0.R/AhT1KWaE9AVNO0OuEY6av/NLYFHiObWN3Fs5oRuQaeZC', '100.0', '2021-12-01 00:15:00-04');
INSERT INTO accounts (id, name, cpf, secret, balance, created_at) VALUES (1005, 'Jimmy Page', '62230356089', '$2a$15$qlHpA/QvOLwcrJ1m0NA6eu8phqJ6zWinW1wArF9eph9K2uZMpsTVK', '666.666', '2022-01-02 14:30:00-04');
INSERT INTO accounts (id, name, cpf, secret, balance, created_at) VALUES (1006, 'Angus Young', '73390504001', '$2a$15$tJHEa9RtE20dIVCzGod5seZ4wViJxzy49a4Ftepc.nPEx0A1FYoNO', '333.333', '2022-03-15 15:19:03-04');
INSERT INTO transfers (id, account_origin_id, account_destination_id, amount,  created_at) VALUES (90000, 1003, 1001, '300.00', '2022-03-15 15:19:03-04');
