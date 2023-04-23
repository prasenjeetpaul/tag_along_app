import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
    selector: 'app-new-group',
    templateUrl: './new-group.component.html',
})
export class NewGroupComponent implements OnInit {

    constructor(private readonly http: HttpClient) { }

    ngOnInit(): void {
        // this.http.get('/users').subscribe(data => console.log(data))
    }

}
