import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
    selector: 'app-new-event',
    templateUrl: './new-event.component.html',
})
export class NewEventComponent implements OnInit {

    constructor(private readonly http: HttpClient) { }

    ngOnInit(): void {
        // this.http.get('/users').subscribe(data => console.log(data))
    }

}
