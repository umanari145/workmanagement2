
-- +migrate Up
INSERT INTO rooms(room_name,address,type,status,note,created,modified) VALUES('市川:A','','10','10','',NOW(),NOW());
INSERT INTO rooms(room_name,address,type,status,note,created,modified) VALUES('市川:B','','10','10','',NOW(),NOW());
INSERT INTO rooms(room_name,address,type,status,note,created,modified) VALUES('市川:C','','10','10','',NOW(),NOW());
INSERT INTO rooms(room_name,address,type,status,note,created,modified) VALUES('本八幡','','10','10','',NOW(),NOW());
INSERT INTO rooms(room_name,address,type,status,note,created,modified) VALUES('自宅','','20','10','',NOW(),NOW());
INSERT INTO rooms(room_name,address,type,status,note,created,modified) VALUES('出勤(外部用)','','10','10','外部スタッフ用出勤',NOW(),NOW());
INSERT INTO rooms(room_name,address,type,status,note,created,modified) VALUES('市川:K','','10','10','',NOW(),NOW());
INSERT INTO rooms(room_name,address,type,status,note,created,modified) VALUES('福岡:A','','10','10','',NOW(),NOW());
INSERT INTO rooms(room_name,address,type,status,note,created,modified) VALUES('福岡:B','','10','10','',NOW(),NOW());
-- +migrate Down
DELETE FROM rooms;