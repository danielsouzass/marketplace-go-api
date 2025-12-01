INSERT INTO payment_methods (key, name) VALUES
  ('boleto', 'Boleto'),
  ('pix', 'Pix'),
  ('cash', 'Dinheiro'),
  ('deposit', 'Depósito Bancário'),
  ('card', 'Cartão de Crédito')
ON CONFLICT (key) DO NOTHING;