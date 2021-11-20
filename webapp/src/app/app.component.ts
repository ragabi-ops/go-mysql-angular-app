import {HttpClient,HttpHeaders, HttpParams, JsonpClientBackend} from '@angular/common/http';
import {Component, ViewChild, AfterViewInit, CUSTOM_ELEMENTS_SCHEMA} from '@angular/core';
import {MatPaginator} from '@angular/material/paginator';
import {MatSort, SortDirection} from '@angular/material/sort';
import {merge, Observable, of as observableOf} from 'rxjs';
import {catchError, map, startWith, switchMap} from 'rxjs/operators';
import { MatTable, MatTableModule } from '@angular/material/table'


/**
 * @title Table retrieving data through HTTP
 */
@Component({
  selector: 'app-root',
  styleUrls: ['app.component.css'],
  templateUrl: 'app.component.html',
})
export class AppComponent implements AfterViewInit {
  displayedColumns: string[] = ['film_id', 'title', 'description', 'rental_duration','rental_rate', 'release_year' ];
  exampleDatabase!: ExampleHttpDatabase | null;
  data: Film[] = [];


  resultsLength = 0;
  isLoadingResults = true;
  isRateLimitReached = false;

  @ViewChild(MatPaginator)
  paginator!: MatPaginator;
  @ViewChild(MatSort)
  sort!: MatSort;

  constructor(private _httpClient: HttpClient) {}

  ngAfterViewInit() {
    this.exampleDatabase = new ExampleHttpDatabase(this._httpClient);

    // If the user changes the sort order, reset back to the first page.
    this.sort.sortChange.subscribe(() => this.paginator.pageIndex = 0);
    function delay(ms: number) {
      return new Promise( resolve => setTimeout(resolve, ms) );
    }
    merge(this.sort.sortChange, this.paginator.page)
      .pipe(
        startWith({}),
        switchMap(() => {
          this.isLoadingResults = true;
          return this.exampleDatabase!.GetAllFilms(
              this.sort.active, this.sort.direction, this.paginator.pageIndex)
            .pipe(catchError(() => observableOf(null)));
        }),
        map(data => {
          // Flip flag to show that loading has finished.
          this.isLoadingResults = false;
          this.isRateLimitReached = data === null;

          if (data === null) {
            return [];
          }
          // Only refresh the result length if there is new data. In case of rate
          // limit errors, we do not want to reset the paginator to zero, as that
          // would prevent users from re-triggering requests.
          this.resultsLength = data.total_count
          return data.items;
        })
      ).subscribe(data => this.data = data);
  }
}

export interface FilmAPI {
  items: Film[];
  total_count: number;
}

export interface Request {
  sort: string
  sort_order: string
  page: number
}
export interface Film {
   film_id: null,
   tite: string,
   descrioption: string,
   release_year: string,
   language_id: number,
   rental_duration: string,
   rental_rate: number,
   length: string,
   replacement_cost: number,
   rating: string,
   last_update: string

}

/** An example database that the data source uses to retrieve data for the table. */
export class ExampleHttpDatabase {
  constructor(private _httpClient: HttpClient) {}
  readonly headers = new HttpHeaders()
    .set('Content-Type', 'application/json');
  requestUrl: string  = "/films"
  GetAllFilms(sort: string, order: SortDirection, page: number): Observable<any> {
    

    if (sort === undefined){
      sort = 'film_id'
    }
    // console.log(`sort=${sort}&sort_order=${order}&page=${page}`)
    
    return this._httpClient.get<any>(`${this.requestUrl}?sort=${sort}&sort_order=${order}&page=${page}`);
  }
}
