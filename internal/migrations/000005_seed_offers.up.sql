INSERT INTO offers (flight_no, airline, departure_airport, arrival_airport, departure_at, arrival_at, currency, price_tiyin, seats_left)
VALUES
    ('HY187','Uzbekistan Airways','TAS','SVO', NOW()+INTERVAL '1 day', NOW()+INTERVAL '1 day 3 hours','UZS',12500000,5),
    ('FE224','SpaceX','TAS','DXB', NOW()+INTERVAL '3 day', NOW()+INTERVAL '4 day 3 hours','UZS',12500000,5),
    ('SU1983','Aeroflot','TAS','SVO', NOW()+INTERVAL '1 day 6 hours', NOW()+INTERVAL '1 day 9 hours','UZS',9800000,2),

    ('TK372','Turkish Airlines','TAS','IST', NOW()-INTERVAL '10 day', NOW()-INTERVAL '10 day' + INTERVAL '4 hours','UZS',8900000,7),
    ('EK201','Emirates','TAS','DXB', NOW()-INTERVAL '7 day', NOW()-INTERVAL '7 day' + INTERVAL '3 hours 30 minutes','UZS',11250000,4),
    ('KC127','Air Astana','TAS','ALA', NOW()-INTERVAL '3 day', NOW()-INTERVAL '3 day' + INTERVAL '1 hour 25 minutes','UZS',5200000,9),
    ('QR583','Qatar Airways','TAS','DOH', NOW()-INTERVAL '1 day', NOW()-INTERVAL '1 day' + INTERVAL '3 hours 45 minutes','UZS',10400000,6),
    ('LH645','Lufthansa','TAS','FRA', NOW()-INTERVAL '15 day', NOW()-INTERVAL '15 day' + INTERVAL '6 hours 10 minutes','UZS',15200000,3),
    ('AF1081','Air France','TAS','CDG', NOW()-INTERVAL '20 day', NOW()-INTERVAL '20 day' + INTERVAL '6 hours 20 minutes','UZS',14750000,2),

    ('S71543','S7 Airlines','TAS','DME', NOW()+INTERVAL '2 hours', NOW()+INTERVAL '5 hours','UZS',7300000,8),
    ('FZ194','flydubai','TAS','DXB', NOW()+INTERVAL '6 hours', NOW()+INTERVAL '9 hours','UZS',9100000,10),
    ('5W701','Wizz Air Abu Dhabi','TAS','AUH', NOW()+INTERVAL '4 hours', NOW()+INTERVAL '7 hours 5 minutes','UZS',6750000,12),
    ('PC7421','Pegasus','TAS','SAW', NOW()+INTERVAL '8 hours', NOW()+INTERVAL '11 hours 10 minutes','UZS',6400000,6),

    ('BT742','airBaltic','TAS','RIX', NOW()+INTERVAL '1 day 2 hours', NOW()+INTERVAL '1 day 7 hours','UZS',8350000,5),
    ('KE992','Korean Air','TAS','ICN', NOW()+INTERVAL '2 day', NOW()+INTERVAL '2 day 6 hours 30 minutes','UZS',19800000,4),
    ('EY256','Etihad','TAS','AUH', NOW()+INTERVAL '2 day 12 hours', NOW()+INTERVAL '2 day 15 hours 30 minutes','UZS',11800000,9),
    ('BA233','British Airways','TAS','LHR', NOW()+INTERVAL '3 day', NOW()+INTERVAL '3 day 7 hours 20 minutes','UZS',17600000,3),
    ('LO652','LOT Polish Airlines','TAS','WAW', NOW()+INTERVAL '4 day', NOW()+INTERVAL '4 day 6 hours','UZS',12100000,7),
    ('AI998','Air India','TAS','DEL', NOW()+INTERVAL '5 day', NOW()+INTERVAL '5 day 2 hours 30 minutes','UZS',5600000,14),
    ('OS848','Austrian Airlines','TAS','VIE', NOW()+INTERVAL '6 day', NOW()+INTERVAL '6 day 6 hours 10 minutes','UZS',13500000,6),
    ('LX257','SWISS','TAS','ZRH', NOW()+INTERVAL '7 day', NOW()+INTERVAL '7 day 6 hours','UZS',14200000,5),

    ('KL461','KLM','TAS','AMS', NOW()+INTERVAL '10 day', NOW()+INTERVAL '10 day 6 hours 15 minutes','UZS',14900000,9),
    ('AZ721','ITA Airways','TAS','FCO', NOW()+INTERVAL '14 day', NOW()+INTERVAL '14 day 6 hours','UZS',13200000,8),
    ('HY188','Uzbekistan Airways','TAS','SVO', NOW()+INTERVAL '12 day', NOW()+INTERVAL '12 day 3 hours','UZS',12750000,11),
    ('EK203','Emirates','TAS','DXB', NOW()+INTERVAL '21 day', NOW()+INTERVAL '21 day 3 hours 20 minutes','UZS',11600000,15),
    ('KC128','Air Astana','TAS','ALA', NOW()+INTERVAL '18 day', NOW()+INTERVAL '18 day 1 hour 25 minutes','UZS',5450000,13),
    ('QR584','Qatar Airways','TAS','DOH', NOW()+INTERVAL '25 day', NOW()+INTERVAL '25 day 3 hours 45 minutes','UZS',10600000,9),
    ('LH646','Lufthansa','TAS','MUC', NOW()+INTERVAL '28 day', NOW()+INTERVAL '28 day 6 hours','UZS',15100000,4),
    ('TK373','Turkish Airlines','TAS','IST', NOW()+INTERVAL '30 day', NOW()+INTERVAL '30 day 4 hours 5 minutes','UZS',9050000,7);
