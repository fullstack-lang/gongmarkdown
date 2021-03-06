// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { ElementDB } from './element-db';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class ElementService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  ElementServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private elementsUrl: string

  constructor(
    private http: HttpClient,
    private location: Location,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.elementsUrl = origin + '/api/github.com/fullstack-lang/gongmarkdown/go/v1/elements';
  }

  /** GET elements from the server */
  getElements(): Observable<ElementDB[]> {
    return this.http.get<ElementDB[]>(this.elementsUrl)
      .pipe(
        tap(_ => this.log('fetched elements')),
        catchError(this.handleError<ElementDB[]>('getElements', []))
      );
  }

  /** GET element by id. Will 404 if id not found */
  getElement(id: number): Observable<ElementDB> {
    const url = `${this.elementsUrl}/${id}`;
    return this.http.get<ElementDB>(url).pipe(
      tap(_ => this.log(`fetched element id=${id}`)),
      catchError(this.handleError<ElementDB>(`getElement id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new element to the server */
  postElement(elementdb: ElementDB): Observable<ElementDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    elementdb.SubElements = []
    elementdb.Rows = []
    let _Element_SubElements_reverse = elementdb.Element_SubElements_reverse
    elementdb.Element_SubElements_reverse = new ElementDB

    return this.http.post<ElementDB>(this.elementsUrl, elementdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        elementdb.Element_SubElements_reverse = _Element_SubElements_reverse
        this.log(`posted elementdb id=${elementdb.ID}`)
      }),
      catchError(this.handleError<ElementDB>('postElement'))
    );
  }

  /** DELETE: delete the elementdb from the server */
  deleteElement(elementdb: ElementDB | number): Observable<ElementDB> {
    const id = typeof elementdb === 'number' ? elementdb : elementdb.ID;
    const url = `${this.elementsUrl}/${id}`;

    return this.http.delete<ElementDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted elementdb id=${id}`)),
      catchError(this.handleError<ElementDB>('deleteElement'))
    );
  }

  /** PUT: update the elementdb on the server */
  updateElement(elementdb: ElementDB): Observable<ElementDB> {
    const id = typeof elementdb === 'number' ? elementdb : elementdb.ID;
    const url = `${this.elementsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    elementdb.SubElements = []
    elementdb.Rows = []
    let _Element_SubElements_reverse = elementdb.Element_SubElements_reverse
    elementdb.Element_SubElements_reverse = new ElementDB

    return this.http.put<ElementDB>(url, elementdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        elementdb.Element_SubElements_reverse = _Element_SubElements_reverse
        this.log(`updated elementdb id=${elementdb.ID}`)
      }),
      catchError(this.handleError<ElementDB>('updateElement'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {

  }
}
