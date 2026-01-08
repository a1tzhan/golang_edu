ALTER TABLE schedule ADD COLUMN room_number VARCHAR(10);

ALTER TABLE schedule DROP COLUMN time_slot;