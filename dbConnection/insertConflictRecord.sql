INSERT INTO tg_users (biba_size, user_name, id_tg_user) VALUES (?, ?, ?)
    ON CONFLICT (id_tg_user) DO UPDATE SET biba_size = EXCLUDED.biba_size, user_name = EXCLUDED.user_name;
