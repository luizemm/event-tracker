CREATE TABLE IF NOT EXISTS event (
	id serial primary key,
	event_type text not null,
	data text not null,
	timestamp text not null
);