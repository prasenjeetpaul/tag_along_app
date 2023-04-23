import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
    selector: 'app-dashboard',
    templateUrl: './dashboard.component.html',
})
export class DashboardComponent implements OnInit {

    constructor(private readonly http: HttpClient) { }

    ngOnInit(): void {
        // this.http.get('/users').subscribe(data => console.log(data))
    }

}
