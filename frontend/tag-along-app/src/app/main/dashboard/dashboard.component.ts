import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { AppService } from 'src/app/app.service';
import { Group } from 'src/app/models/group.model';
import Swal from 'sweetalert2'

@Component({
    selector: 'app-dashboard',
    templateUrl: './dashboard.component.html',
})
export class DashboardComponent implements OnInit {

    groups: Group[] = [];

    constructor(
        private readonly http: HttpClient,
        public readonly appService: AppService
    ) { }

    ngOnInit(): void {
        this.http.get<Group[]>('/groups/user/' + this.appService.loggedInUser?.id)
            .subscribe({
                next: data => this.groups = data,
                error: err => Swal.fire({
                    icon: 'error',
                    title: 'Internal Server Error',
                    text: (err && err.error) ? err.error.error : 'Something went wrong fetching group details',
                })
            });
    }

}
