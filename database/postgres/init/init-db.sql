CREATE TABLE segments(
    id SERIAL NOT NULL,
    name text NOT NULL,
    PRIMARY KEY(id)
);
CREATE UNIQUE INDEX segments_name_key ON "segments" USING btree ("name");
CREATE TABLE user_segments(
    user_id bigint NOT NULL,
    segment_id bigint NOT NULL,
    PRIMARY KEY(user_id,segment_id),
    CONSTRAINT fk_user_segments_segment FOREIGN key(segment_id) REFERENCES segments(id)
);