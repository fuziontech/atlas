-- atlas:delimiter \n-- end --\n

DELIMITER $$
CREATE DEFINER='boring' PROCEDURE proc ()
    COMMENT 'ATLAS_DELIMITER'
    SQL SECURITY INVOKER
    NOT DETERMINISTIC
    MODIFIES SQL DATA
BEGIN
    UPDATE performance_schema.threads
    SET instrumented = 'YES'
    WHERE type = 'BACKGROUND';

    SELECT CONCAT('Enabled ', @rows := ROW_COUNT(), ' background thread', IF(@rows != 1, 's', '')) AS summary;
END$$
DELIMITER ;

-- end --

CALL proc();