# Contributing to ani

Thank you for your interest in contributing to ani! We follow a strict development workflow to ensure quality and consistency. Please read the rules below carefully before making any changes.

## Development Workflow

### 1. Branch First
Always create a new branch for any feature addition or modification. Do not commit directly to the main branch.

```bash
git checkout -b feature/your-feature-name
```

### 2. Documentation Driven Development
We adhere to a "Documentation First" approach. Before writing any code, you must update the project documentation in the following order:

1.  **Update `requirements.md`**: Discuss and clarify requirements with the user/stakeholders. Document the agreed-upon requirements.
2.  **Update `design.md`**: Detail the architectural changes or design decisions needed to satisfy the requirements.
3.  **Update `tasks.md`**: Break down the implementation into granular, trackable tasks.
4.  **Implement**: Write the code to fulfill the tasks.
5.  **Update `README.md`**: If the changes affect how the user interacts with the software (e.g., new commands, configuration changes, UI updates), you MUST update the `README.md` to reflect these changes.

## Pull Requests
Ensure your PR description links to the relevant items in `tasks.md`.
