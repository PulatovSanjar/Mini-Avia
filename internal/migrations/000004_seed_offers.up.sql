INSERT INTO offers (flight_no, airline, departure_airport, arrival_airport, departure_at, arrival_at, currency, price_tiyin, seats_left)
VALUES
    ('HY187','Uzbekistan Airways','TAS','SVO', NOW()+INTERVAL '1 day', NOW()+INTERVAL '1 day 3 hours','UZS',12500000,5),
    ('FE224','SpaceX','TAS','DXB', NOW()+INTERVAL '3 day', NOW()+INTERVAL '4 day 3 hours','UZS',12500000,5),
    ('SU1983','Aeroflot','TAS','SVO', NOW()+INTERVAL '1 day 6 hours', NOW()+INTERVAL '1 day 9 hours','UZS',9800000,2);
