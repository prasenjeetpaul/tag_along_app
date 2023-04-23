import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ActivatedRoute, Router } from '@angular/router';
import { Group } from 'src/app/models/group.model';
import { concatMap, delay } from 'rxjs';
import Swal from 'sweetalert2';

@Component({
    selector: 'app-group-detail',
    templateUrl: './group-detail.component.html',
})
export class GroupDetailComponent implements OnInit {

    groupData!: Group
    constructor(
        private readonly http: HttpClient,
        private readonly route: ActivatedRoute,
        private readonly router: Router
    ) { }

    ngOnInit(): void {
        window.scrollTo({ top: 0, behavior: 'smooth' });
        this.route.params.pipe(concatMap(params => this.http.get<Group>('/groups/' + params['id'])))
            .subscribe({
                next: data => this.groupData = data,
                error: _ => {
                    Swal.fire({
                        icon: 'error',
                        title: 'Server Error',
                        text: 'Something went wrong! Please try again'
                    })
                    this.router.navigate(['/']);
                }
            })
    }

}
