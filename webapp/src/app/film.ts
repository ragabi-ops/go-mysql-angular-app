export class Film {
    constructor(
        public film_id?: null,
        public tite?: string,
        public descrioption?: string,
        public release_year?: string,
        public language_id?: number,
        public rental_duration?: string,
        public rental_rate?: number,
        public length?: string,
        public replacement_cost?: number,
        public rating?: string,
        public last_update?: string

      ) {}
}