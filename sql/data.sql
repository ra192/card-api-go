INSERT INTO public.merchants (id, name, secret) VALUES (default , 'Internal', 'C5X+rivjAkiczaDqkDwMKAVxN+cSGrudaGCRhWZeVos=');
INSERT INTO public.merchants (id, name, secret) VALUES (default , 'Wayne Enterprise', '/05W5ibIDcNPkHXS8mC71CKB+10Ct9gzJ41C9HRMYeQ=');

INSERT INTO public.accounts (id, active, currency, name, merchant_id) VALUES (default , true, 'USD', 'Cash account', 1);
INSERT INTO public.accounts (id, active, currency, name, merchant_id) VALUES (default , true, 'USD', 'Card account', 1);
INSERT INTO public.accounts (id, active, currency, name, merchant_id) VALUES (default , true, 'USD', 'Fee account', 1);
INSERT INTO public.accounts (id, active, currency, name, merchant_id) VALUES (default , true, 'USD', 'Wayne USD account', 2);