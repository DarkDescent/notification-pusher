-- postgres SQL

create table if not exists firebase_token (
    id bigserial primary key,
    token text not null,
    cid uuid not null,
    active boolean not null,
    expiresAt timestamp not null
);

create index if not exists "firebase_token_cid_token_idx" on firebase_token(cid, token);

-- comment on table firebase_token is 'Table that holds all tokens for different users that can get notifications in their mobile apps. We make it dependent on firebase because primary use case of our application on today - use only FCM support. Other mechanisms can use other API contracts and other data.';

-- comment on column firebase_token.token is 'Token that we use to send information through firebase messaging. By using this token we can send notification to specific mobile app instance.'
-- comment on column firebase_token.cid is 'Identifier of customer in some "internal system". We expect that somewhere (not in notification service) we store all information about customers. But this service contains only information about how to notify specific customer.';
-- comment on column firebase_token.active is 'Is our firebase token active? We expect that it can expires. Or can be invalid according to FCM specification.';
-- comment on column firebase_token.expiresAt is 'Our token can live only some time. And must be updated periodically from mobile app.';

