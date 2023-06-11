type Project = {
	name: string;
	href: string;
	description: string;
	tags: string[];
};

export async function getProjects(noOfItems?: number, page?: number): Promise<Project[]> {
	const github = await fetch(
		`https://api.github.com/users/raphgl/repos?type=sources&sort=pushed&per_page=${noOfItems}&page=${page}`
	);
	const projectsJSON = await github.json();

	let projects = [];
	for (const project of projectsJSON) {
		if (!project.fork)
			projects.push({
				name: project.name,
				href: project.html_url,
				description: project.description,
				tags: project.topics
			});
	}

	return projects;
}
