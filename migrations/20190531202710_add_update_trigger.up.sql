CREATE FUNCTION notify_new_event() 
RETURNS trigger AS $$  
BEGIN  
	PERFORM pg_notify('event_change', row_to_json(NEW)::text);
	RETURN NEW;
END;  
$$ LANGUAGE plpgsql;

CREATE TRIGGER event_change 
	BEFORE INSERT ON events  
	FOR EACH ROW 
	EXECUTE PROCEDURE notify_new_event();
