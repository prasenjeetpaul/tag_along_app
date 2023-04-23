import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
    selector: 'app-group-detail',
    templateUrl: './group-detail.component.html',
})
export class GroupDetailComponent implements OnInit {

    constructor(private readonly http: HttpClient) { }

    ngOnInit(): void {
        // this.http.get('/users').subscribe(data => console.log(data))
    }

}
