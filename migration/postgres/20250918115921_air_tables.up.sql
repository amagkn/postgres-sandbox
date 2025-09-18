-- таблица компаний
CREATE TABLE air.company
(
    id   INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- таблица пассажиров
CREATE TABLE IF NOT EXISTS air.passenger
(
    id   INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- таблица рейсов
CREATE TABLE IF NOT EXISTS air.trip
(
    id         INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    company_id INT NOT NULL,
    plane      VARCHAR(255),
    town_from  VARCHAR(255),
    town_to    VARCHAR(255),
    time_out   TIMESTAMP WITHOUT TIME ZONE,
    time_in    TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT fk_trip_company FOREIGN KEY (company_id) REFERENCES air.company (id)
);

-- таблица связи пассажиров и рейсов
CREATE TABLE IF NOT EXISTS air.pass_in_trip
(
    id           INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    trip_id      INT NOT NULL,
    passenger_id INT NOT NULL,
    place        VARCHAR(50),
    CONSTRAINT fk_pass_in_trip_trip FOREIGN KEY (trip_id) REFERENCES air.trip (id),
    CONSTRAINT fk_pass_in_trip_passenger FOREIGN KEY (passenger_id) REFERENCES air.passenger (id)
);

-- индексы для ускорения поиска по внешним ключам
CREATE INDEX IF NOT EXISTS idx_trip_company ON air.trip(company_id);
CREATE INDEX IF NOT EXISTS idx_pass_in_trip_trip ON air.pass_in_trip(trip_id);
CREATE INDEX IF NOT EXISTS idx_pass_in_trip_passenger ON air.pass_in_trip(passenger_id);