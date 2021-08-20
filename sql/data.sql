INSERT INTO public.merchants (id, name, secret) VALUES (1, 'Internal', 'C5X+rivjAkiczaDqkDwMKAVxN+cSGrudaGCRhWZeVos=');
INSERT INTO public.merchants (id, name, secret) VALUES (2, 'Wayne Enterprise', '/05W5ibIDcNPkHXS8mC71CKB+10Ct9gzJ41C9HRMYeQ=');

INSERT INTO public.accounts (id, active, currency, name, merchant_id) VALUES (1, true, 'USD', 'Cash account', 1);
INSERT INTO public.accounts (id, active, currency, name, merchant_id) VALUES (2, true, 'USD', 'Card account', 1);
INSERT INTO public.accounts (id, active, currency, name, merchant_id) VALUES (3, true, 'USD', 'Fee account', 1);
INSERT INTO public.accounts (id, active, currency, name, merchant_id) VALUES (4, true, 'USD', 'Wayne USD account', 2);