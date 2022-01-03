use actix_web::web;

pub mod echo;
pub mod api;

pub fn init(cfg: &mut web::ServiceConfig) {
    cfg
    .service(echo::hello)
    .service(echo::echo)
    .service(api::fetch_most_popular_videos)
    .service(api::fetch_related_videos)
    .service(api::search_video);
}