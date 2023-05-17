type Project = {
	name: string;
	href: string;
	description: string;
	tags: string[];
};

export async function getProjects(noOfItems?: number, page?: number): Promise<Project[]> {
	const github = await fetch(
		'https://github.com/RaphGL?tab=repositories&q=&type=source&language=&sort=stargazers'
	);
	const pageText = await github.text();
	const parser = new DOMParser();
	const pageDOM = parser.parseFromString(pageText, 'text/html');

	let projectsElems = pageDOM.querySelectorAll('div.col-10.col-lg-9.d-inline-block');
	let projects: Project[] = [];

	for (let elem of projectsElems) {
		let projectName = elem.querySelector('h3.wb-break-all a');
		let projectDescription = elem.querySelector('p.col-9.d-inline-block.color-fg-muted.mb-2.pr-4');
		let projectTagElems = elem.querySelectorAll('a.topic-tag.topic-tag-link.f6.my-1');

		let projectTags = [];
		for (let tag of projectTagElems) {
			projectTags.push(tag.innerHTML);
		}

		console.log(projectTags);

		projects.push({
			name: projectName?.innerHTML as string,
			href: ('https://github.com' + projectName?.getAttribute('href')) as string,
			description: projectDescription?.innerHTML as string,
			tags: projectTags
		});
	}

	if (noOfItems) {
		return projects.slice(0, noOfItems);
	} else if (noOfItems && page) {
		return projects.slice(page * noOfItems, noOfItems);
	} else {
		return projects;
	}
}
