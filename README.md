# 12 Week Go Learning Plan
version of go: 1.27.1

Slammed together with chatgpt for what i want and tweaked by myself.
I have a computer science degree but i'm dumb as fuck so i want to actually learn.

**Outcome Wanted:** Build, test, explain and finish Go programs then create a capstone project

## Curriculum

- Each week is "learning stage". ~30hour schedule
- Projects should be done in order 1-12. Only skip a repeated exercise if I can complete a check without help
- Early projects should be doable in one focused session. Laters ones can take several. It's okay if setup or debugging adds time
- Time spent on the capstone comes out weekly budget. Not extra work
- Option extensions and bonus projects are a menu. Ensure core behaviour before choosing one.
- **KEEP APPLYING TO JOBS WHILE DOING THIS. PLEASE OFFER ME A JOB. PLEASE OFFER ME A JOB. PLEASE PLEASE PLEASE**

### Daily Plan

1. **Recall — 10 minutes:** Explain yesterday's idea without checking code
2. **Read — 20-30 mintes:** Read documentaion for the day's feature.
3. **Build — 30-120 minutes:** implement a feature
4. **Check — 20-30 minutes:** test a normal input, boundary input and invalid input
5. **Record — 10 minutes:** commit a working change and note the concept learned

Just a guidance, dunno if i will do more or less but i'll try to stick to them.

### Example weekly rhytm

| Day | Task                                                                            |
| --- | ------------------------------------------------------------------------------- |
| 1   | Learn the weeks main concepts and start its project                             |
| 2   | Finish it and begin the small build                                             |
| 3   | Work through the next project or a difficult concept.                           |
| 4   | Complete the main weekly deliverable or integrate it into DataWatch.            |
| 5   | Review, fix a bug, demonstrate the result and attempt one small change unaided. |
| 6–7 | Optional practice, applications, catching up or time off.                       |

There are fewer new projects in the later weeks because integration, reliability and explanation take longer.

## Repo setup

Everything stored in this repo with one root Go module. Each executable will have its own directory; Two unrelated main functions cannot share a package directory. A repository with many exercises does not require a go.mod file in every exercise. [R2]

| Path               | Purpose                                       |
| ------------------ | --------------------------------------------- |
| README.md          | Roadmap, progress table.                      |
| go.mod             | One module for all learning exercises.        |
| week01/converter/  | Project 01 and its content.                   |
| week01/calculator/ | Project 02 and its tests.                     |
| weekX/dirY/        | Project Z                                     |
| notes/week01.md    | What I learned, what failed what to practice. |

Tests beside the code they test. Broken experiments go in notes rather than commits.

### What "finished" means

For each project, be able to:

- Run it using a command documented in its README.
- Explain the main types and the route form input to output
- Show the card's completion check passing
- Describe one invalid input and the resulting behaviour
- Make one small change without copying a solution

README only needs to be short: purpose, run command, example I/O, expected failure behaviour and one thing learnt

## The route at a glance

| week | projects | main learning                                  | weekly result                                 |
| ---- | -------- | ---------------------------------------------- | --------------------------------------------- |
| 1    | 01–04    | syntax, functions, control flow, slices, tests | a useful statistics cli.                      |
| 2    | 05–08    | maps, structs, methods, pointers, errors       | several small programs you can explain.       |
| 3    | 09–12    | files, json, csv, readers, writers             | a tested csv profiler.                        |
| 4    | 13–16    | http clients, handlers, local http tests       | datawatch accepts a csv over http.            |
| 5    | 17–20    | sql, migrations, transactions, history         | reports survive restarts.                     |
| 6    | 21–24    | configuration, quality rules, comparisons      | first useful local datawatch release.         |
| 7    | 25–27    | goroutines, channels, limits, cancellation     | a controlled batch processor.                 |
| 8    | 28–30    | job states, durable work, recovery             | background processing with restart behaviour. |
| 9    | 31–33    | html templates, forms, usable results          | a browser-based demonstration.                |
| 10   | 34–36    | configuration, logs, shutdown, integration     | a reproducible and diagnosable application.   |
| 11   | 37–38    | benchmarks, profiling, fuzzing                 | evidence about correctness and performance.   |
| 12   | 39–40    | automation, api clients, release discipline    | a demonstrable portfolio release.             |

## Week 1 — Become comfortable writing go

**Read first:** Basic syntax and flow-control sections of A Tour of Go. Learn packages, imports, variables, types, functions, multiple return values and slices as the projects require them. [R1]

### Project 01. Unit Converter

**Requires:** No prior Go. **Practise:** Variables, arithmetic, functions, formatted output.
**Build:** CLI program that converts Celsius to Fahrenheit and kilometres to miles.
**Done when:** 0°C produces 32°F, negative temperatures work, and invalid input produces a useful message.

### 02. Four operation calculator

**Time:** 2-3hours. **Requires:** 01. **Practise:** conditionals or switch, return values, errors, basic testing.
**Build:** a program that takes two numbers and an operation

1. Implement addition and subtraction as functions.
2. Add multiplication and division
3. Make division by zero and unsupported operations return errors.
4. Keep input parsing and printing outside the calculation functions
5. Add test for normal artihmetic, negative numbers and division by zero

**Done when:** 8 divided by 2 gives 4; dividing by zero does not panic; invalid text is rejected rather than treated as zero
**Option extension:** allow previous result to become the next operation's first operand.

### 03. Number-guessing game

**Time:** 2–3 hours. **Requires:** 01–02. **Practise:** loops, branches, state, random values.

**Build:** a game that chooses an integer from 1 to 100 and gives higher/lower hints.

1. Begin with a fixed secret so behaviour is easy to check.
2. Read guesses repeatedly and give a hint.
3. Count valid attempts and handle guesses outside the allowed range.
4. Add a quit command and ensure end-of-input exits cleanly.
5. Replace the fixed secret with a random one; keep the comparison logic testable using a chosen secret.

**Done when:** a player can win, quit or enter bad text without breaking the game; the attempt count follows a documented rule.

**Optional extension:** add difficulty levels with different ranges. Use math/rand/v2 for game randomness; it is not a security-token generator. [R13]

### 04. Number statistics CLI

**Time:** 3–4 hours. **Requires:** 02. **Practise:** slices, loops, accumulation, argument parsing, testing.

**Build:** a tool that reports count, sum, minimum, maximum and mean for supplied numbers.

1. Write functions over an existing slice of numbers before handling terminal arguments.
2. Initialise minimum and maximum from actual data, not from an assumed zero.
3. Define what an empty input means; return an error or an explicitly empty result.
4. Parse command-line arguments and keep parsing separate from statistics.
5. Test positive, negative, decimal, single-element, empty and invalid inputs. Compare floating-point results with a justified tolerance where needed.

**Done when:** 2, 4, 6 gives count 3, sum 12 and mean 4; all-negative input produces the correct maximum.

**Optional extension:** add median and explain whether sorting changes the caller's slice.

**Weekly checkpoint:** write a new “range = maximum minus minimum” feature from its requirement. Explain package versus module, integer versus floating-point division, and why empty input needs attention.

## Week 2 — Organise data and handle mistakes

**Read first:** Tour sections on maps, structs, pointers and methods; the testing documentation when writing table-driven cases. [R1, R3]

### 05. Word-frequency analyser

**Time:** 2–3 hours. **Requires:** 04. **Practise:** strings, maps, slices, sorting.

**Build:** count words and show the most common ones.

1. Define what counts as a word; start with whitespace-separated tokens.
2. Normalise case and count tokens in a map.
3. Convert the counts into sortable results.
4. Sort by frequency, breaking ties alphabetically so output is repeatable.
5. Add a top-N option and test empty input, mixed case and tied counts.

**Done when:** “Go go learn” gives go: 2 and learn: 1, and tied results always appear in the same order.

**Optional extension:** document and implement punctuation handling. Decide whether a character means a byte, a Unicode code point or something else.

### 06. In-memory task tracker

**Time:** 3–4 hours. **Requires:** 05. **Practise:** structs, methods, pointer receivers, identity.

**Build:** add, list, complete and remove tasks during one program run.

1. Define a task with an ID, title and completion state.
2. Store tasks in a slice or map; explain your choice.
3. Add operations for creating and locating a task.
4. Add completion and removal, returning useful errors for unknown IDs.
5. Put an interactive menu around the operations and test state changes separately.

**Done when:** removing task 2 does not make task 3 change its ID, and completing an unknown task is handled deliberately.

**Optional extension:** filter open versus completed tasks. Persistence comes in week 3.

### 07. Terminal quiz

**Time:** 2–3 hours. **Requires:** 06. **Practise:** slices of structs, string handling, small functions.

**Build:** a five-question quiz with a final score.

1. Represent each question, answer and explanation as data.
2. Ask questions one at a time and collect the answer.
3. Decide whether comparison ignores surrounding spaces or case.
4. Count correct answers and show explanations for mistakes.
5. Separate answer checking from terminal input, then test it.

**Done when:** every question is asked once, the score is correct, and an empty answer has defined behaviour.

**Optional extension:** shuffle question order while keeping each answer attached to its question.

### 08. Service-log summariser

**Time:** 3–5 hours. **Requires:** 04–06. **Practise:** parsing, structs, errors, grouped statistics.

**Build:** summarise records such as “api 200 12.5”, representing service, status and duration in milliseconds.

1. Document the three-field format and define a record type.
2. Write a parser that returns a record or an explanatory error.
3. Aggregate request count, status codes of 400 or above, and mean duration by service.
4. Start with lines already in memory; choose whether a bad line stops processing or is reported and skipped.
5. Add table-driven tests for missing fields, bad numbers, negative durations and multiple services.

**Done when:** api records with durations 10 and 30 have a mean of 20; malformed records cannot silently become successful requests.

**Optional extension:** show the slowest request per service.

**Weekly checkpoint:** explain a zero value, a pointer receiver and an error return. Rebuild one parser from its description. Add a small CI workflow for formatting checks, tests and go vet if you have connected the repository to CI.

## Week 3 — Make useful tools with files

**Read first:** io.Reader/io.Writer, file handling, defer, encoding/json and encoding/csv. Use the CSV parser for quoted fields; splitting a line on commas is insufficient. [R4, R5]

### 09. JSON-backed notes CLI

**Time:** 3–4 hours. **Requires:** 06. **Practise:** files, JSON, flags, persistence, errors.

**Build:** add, list and remove notes that survive a restart.

1. Define a note and a JSON file format.
2. Implement load and save functions independently of the command-line interface.
3. Add add/list/remove commands, or adapt project 06 into a persistent task tracker.
4. Treat a missing file as an empty first run; treat corrupt JSON as an error.
5. Save via a temporary file in the same directory, close it, and replace the destination only after writing succeeds; handle each error.

**Done when:** notes survive a new process, and a corrupt file is reported without being silently overwritten.

**Optional extension:** export notes as plain text. Keep this exercise single-process; temporary-file replacement is not a complete concurrency or crash-durability strategy.

### 10. CSV expense analyser

**Time:** 3–4 hours. **Requires:** 04 and 09. **Practise:** CSV records, dates, grouping, precise amounts.

**Build:** total expenses by category and month.

1. Create a tiny fixture with date, category and amount columns.
2. Parse records using encoding/csv and dates with an explicit layout.
3. Parse pounds and pence into integer pence; reject unsupported precision instead of relying on binary floating-point money totals.
4. Group totals by category and month.
5. Report bad rows with their record number and test quoted commas, invalid dates and empty files.

**Done when:** £0.10 plus £0.20 totals exactly £0.30; a quoted category containing a comma remains one field.

**Optional extension:** write a summary CSV that spreadsheet software can open.

### 11. Duplicate-file report

**Time:** 3–5 hours. **Requires:** 09. **Practise:** directory traversal, file metadata, streaming, hashing.

**Build:** report files with identical contents within a directory.

1. Walk regular files in a small test directory and define how symbolic links are handled.
2. Group candidates by size before reading contents.
3. Compute SHA-256 using streamed reads for candidate groups.
4. Group matching hashes and, if claiming exact identity, compare candidate bytes before reporting a duplicate.
5. Sort paths and report unreadable files clearly. The core tool produces a report without changing files.

**Done when:** two differently named copies are grouped, and same-sized files with different contents are not.

**Optional extension:** estimate potentially recoverable space, counting one retained copy per group.

### 12. CSV profiler

**Time:** 4–6 hours. **Requires:** 04, 08 and 10. **Practise:** io.Reader, incremental parsing, explicit data semantics.

**Build:** the first small piece of DataWatch: profile a CSV's columns and rows.

1. Define supported input: UTF-8, comma separator, one header row and a policy for duplicate headers; reject invalid UTF-8 fields deliberately.
2. For each column, calculate missing count and the count, minimum, maximum and mean of successfully parsed finite numeric values.
3. Define missingness, whitespace handling and invalid numeric values. Represent an unavailable numeric summary explicitly rather than as zero.
4. Process records incrementally and produce a structured report plus a terminal summary.
5. Test quoted commas, quoted newlines, header-only input, malformed rows and mixed values using tiny fixtures with hand-calculated answers.

**Done when:** a column containing 10, an empty cell, 20 and “bad” reports one missing value, two numeric values and numeric mean 15.

**Optional extension:** JSON output. Exact distinct counts can use growing memory; they are a separate feature.

**Weekly checkpoint:** run all four tools on fixture files. Explain why an io.Reader lets the profiler process files and in-memory test input, and why an empty numeric population must not be shown as a measured zero.

## Week 4 — Send data over HTTP

**Read first:** net/http for client and server basics; net/http/httptest for local tests. Keep the CSV calculations independent of request handling. [R6, R7]

### 13. Endpoint checker

**Time:** 2–3 hours. **Requires:** 08. **Practise:** HTTP clients, timeouts, response handling.

**Build:** a terminal tool that checks a supplied list of local or public HTTP addresses.

1. Make one request with an http.Client configured with a timeout.
2. Record the address, status code and elapsed time.
3. Close response bodies, including for non-200 responses.
4. Loop over several addresses sequentially; print a summary.
5. Test against httptest servers that return success, errors and a deliberately delayed response.

**Done when:** one slow server cannot block the tool indefinitely, and a 500 response is reported separately from a transport error.

**Optional extension:** accept a custom success-status range.

### 14. JSON API client

**Time:** 3–4 hours. **Requires:** 09 and 13. **Practise:** JSON decoding, response validation, client/server separation.

**Build:** fetch a small JSON document from a local test server and print selected fields.

1. Design a tiny JSON payload, such as a reading list with title and status.
2. Write an httptest server that returns this fixture.
3. Make a client request with a timeout and validate the status code.
4. Decode into a Go struct; handle invalid JSON and a missing required field.
5. Put decoding and display in separate functions and test both normal and malformed responses.

**Done when:** normal JSON prints the expected fields, and a non-2xx response or malformed JSON returns an explanatory error.

**Optional extension:** save a fetched result to a JSON file.

### 15. CSV report server

**Time:** 3–4 hours. **Requires:** 12 and 14. **Practise:** handlers, JSON responses, routing.

**Build:** a local server that exposes a fixed, already-created CSV profile.

1. Create a small router with a health endpoint.
2. Load a known fixture during startup and compute its profile.
3. Return that report as JSON from a GET route.
4. Set the response content type and return controlled errors if the fixture cannot be read.
5. Test the handler with httptest and compare decoded JSON fields.

**Done when:** a browser or HTTP client can fetch the known report, and a bad fixture prevents an apparently successful report.

**Optional extension:** add a second fixture selected by a known ID.

### 16. Upload and profile endpoint

**Time:** 4–6 hours. **Requires:** 12 and 15. **Practise:** request boundaries, handler tests, package separation.

**Build:** DataWatch's first synchronous profiling endpoint. This is the point to start its portfolio repository, if you follow the earlier plan.

1. Document the request: POST a raw CSV body, content type text/csv, and an example response shape.
2. Limit request size, for example to 10 MiB, and reject oversized requests; count actual bytes read rather than trusting a header.
3. Pass the bounded request body to the CSV profiler through an io.Reader.
4. Return a JSON profile on success and a consistent JSON error for unsupported content, malformed CSV and oversized input.
5. Test a clean fixture, a malformed fixture, an oversized body and an empty body using httptest.

**Done when:** the same profiling logic works from a CLI fixture and an HTTP request, with no HTTP-specific logic in the profile package.

**Optional extension:** add a CLI flag that posts a file to the service. Leave browser multipart uploads until week 9.

**Weekly checkpoint:** trace an upload through handler, profiler and response. Save a reproducible curl example and a tiny fixture in DataWatch. This is already a demonstrable Go service.

## Week 5 — Store and query results

**Read first:** SQL basics, constraints and transactions; the Go database guide. Go's database/sql API needs a compatible PostgreSQL driver. Use explicit queries and versioned migrations. [R8, R9]

### 17. PostgreSQL contact book

**Time:** 4–6 hours. **Requires:** 09. **Practise:** connections, parameterised SQL, rows and errors.

**Build:** a focused command-line contact book using a local PostgreSQL instance.

1. Start a local development database and create contacts with an ID, name and email.
2. Write a migration that creates the table, including a constraint preventing duplicate emails.
3. Implement add, get, list and delete using parameterised queries.
4. Handle “not found”, duplicate email and unavailable database as distinct outcomes.
5. Run a test against an isolated test database; restore its schema from the migration.

**Done when:** contacts survive a new process and entering duplicate email does not create a second row.

**Optional extension:** add case-insensitive email uniqueness after deciding how it should work.

### 18. Transactional CSV importer

**Time:** 3–5 hours. **Requires:** 10 and 17. **Practise:** validation before persistence, transactions.

**Build:** import a CSV of contacts, either into project 17 or into a separate test table.

1. Parse the header and rows and reject invalid email or missing name.
2. Decide whether a bad row rejects the entire file or is skipped; start with all-or-nothing.
3. Insert rows in a transaction; commit only when every row succeeds.
4. Show how many records were inserted on success, and report the failing record on error.
5. Test a file whose later row violates a constraint and verify the earlier rows are rolled back.

**Done when:** a failed import leaves the database in its pre-import state.

**Optional extension:** add a dry-run mode that validates and counts without writing.

### 19. Dataset registry

**Time:** 4–6 hours. **Requires:** 12, 16 and 17. **Practise:** schema design, stored reports, service integration.

**Build:** add dataset and report history to DataWatch.

1. Define dataset, upload version and processing-run records with IDs and foreign keys.
2. Store uploaded files under application-generated storage identifiers, not user-supplied paths.
3. Save the profile, input metadata and run status; use a transaction for related database writes.
4. Add an endpoint to list datasets and an endpoint to retrieve a report by ID.
5. Restart both app and database, then retrieve the same report; test missing and invalid IDs.

**Done when:** one uploaded version can be retrieved after a restart, and a failed save cannot appear as a successful complete run.

**Optional extension:** record the input file's SHA-256 in its metadata.

### 20. Run-history browser

**Time:** 3–4 hours. **Requires:** 19. **Practise:** SQL joins, pagination, ordering.

**Build:** an API that lists runs for one dataset.

1. Define a response with run ID, file name, creation time and status.
2. Join the necessary tables in one readable query.
3. Set an explicit newest-first ordering with a stable tie-breaker.
4. Add a limited page size and a way to request the next page.
5. Test one dataset with several versions and a second dataset whose runs must not leak into the first list.

**Done when:** two requests for successive pages contain each matching run once in the expected order.

**Optional extension:** filter by successful versus failed processing.

**Weekly checkpoint:** document database startup, migrations and test setup. Explain a foreign key, a parameterised query, the transaction boundary and why a user-supplied filename is not a storage path.

## Week 6 — Make reports meaningful

**Read first:** JSON configuration, validation errors and rule semantics. A failed quality check is a successful analysis that found bad data; it is different from a parser or job failure.

### 21. Rule-file checker

**Time:** 2–4 hours. **Requires:** 09 and 12. **Practise:** structured configuration and validation.

**Build:** check a JSON file containing quality rules.

1. Choose a small schema: required columns, columns that must be numeric, numeric ranges and maximum missing fractions.
2. Decode JSON into explicit Go types.
3. Validate contradictory or impossible settings, such as a range whose minimum exceeds its maximum, a fraction outside 0–1 or an empty column name.
4. Print each valid rule in a human-readable form.
5. Test a valid file, malformed JSON, unknown rule type and invalid threshold.

**Done when:** invalid configuration is rejected before any CSV is processed, with an error that identifies the offending rule.

**Optional extension:** reject unexpected JSON fields and document the rule schema.

### 22. Column-rule evaluator

**Time:** 4–6 hours. **Requires:** 12 and 21. **Practise:** translating requirements into testable behaviour.

**Build:** evaluate quality rules against records and the profile.

1. Implement required-column and required-value checks.
2. Implement numeric and numeric-range checks, distinguishing empty values from non-numeric values.
3. Implement maximum missing fraction using a documented denominator; define the empty-dataset case.
4. Return structured violations with rule, column, observed value and threshold or row where relevant.
5. Test a clean fixture and a corrupted one with independently calculated expected violations.

**Done when:** an absent required column, one bad number and a high missing fraction produce three distinguishable findings; a clean file produces no findings.

**Optional extension:** report a bounded sample of failing row numbers, rather than retaining every bad row.

### 23. Baseline comparison tool

**Time:** 3–5 hours. **Requires:** 12 and 22. **Practise:** explicit missing data, comparisons.

**Build:** compare two stored reports, with an explicitly chosen baseline.

1. Identify added and removed columns.
2. Compare total rows and missing fractions per common column.
3. Compare numeric count and mean only when both reports have those values.
4. Represent unavailable metrics clearly; express changes in missing fractions as percentage points.
5. Test a baseline with no rows, a removed column and a column that changes from numeric to non-numeric.

**Done when:** the output distinguishes “new value is zero”, “metric is unavailable” and “column no longer exists”.

**Optional extension:** output a compact JSON change list. A mean difference is descriptive, not proof of statistical drift.

### 24. Local quality-gate CLI

**Time:** 4–6 hours. **Requires:** 19, 21–23. **Practise:** assembling several packages into one flow.

**Build:** a terminal flow that accepts a CSV, a rule file and optionally a baseline report; make the same behaviour available in DataWatch.

1. Parse inputs and validate rules before recording a run.
2. Profile the CSV, evaluate rules and save the exact configuration with the resulting run.
3. Produce a readable summary and optional baseline comparison.
4. Give parsing failure, quality failure and success distinct documented outcomes.
5. Write an integration test for one clean CSV and one intentionally corrupted CSV.

**Done when:** a complete local upload, stored report, rule check and comparison works from the documented command or API. Tag a first local DataWatch release.

**Optional extension:** define an exit code useful in automation when quality rules fail.

**Weekly checkpoint:** show the clean and corrupted example end to end. If this flow is incomplete, finish it before starting the background jobs in week 8.

## Week 7 — Concurrency without guesswork

**Read first:** goroutines, channels, sync and context. First make one-file processing correct; then run independent files concurrently. Use go test -race on tests that exercise shared work. [R10, R11]

### 25. Parallel file analyser

**Time:** 4–5 hours. **Requires:** 12. **Practise:** goroutines and collecting results.

**Build:** profile several independent fixture files during one invocation.

1. Start by profiling the list sequentially and recording the result for each file.
2. Start a goroutine for each file in a small, fixed list.
3. Return each result and its filename through a result channel; wait for all expected results.
4. Sort final output by filename so scheduling does not scramble the report.
5. Test both successful and malformed files, then run the exercised tests with the race detector.

**Done when:** every requested file yields exactly one result or error and the program exits on a mixed successful/failed batch.

**Optional extension:** compare timings on real files, keeping I/O, CPU and the test machine in mind.

### 26. Bounded worker-pool processor

**Time:** 4–6 hours. **Requires:** 25. **Practise:** channels, WaitGroup, backpressure, ownership.

**Build:** replace “one goroutine per file” with a small configurable worker limit.

1. Create a queue of file jobs and start a fixed number of workers.
2. Have each worker process one job at a time and send one result.
3. Ensure the producer, workers and result collector can all finish if one file fails.
4. Use an explicit signal in a test to verify the active count never exceeds the limit; avoid timing assertions based on sleep.
5. Run go test -race and explain who closes each channel.

**Done when:** twenty files with a worker limit of three never have more than three active processing jobs.

**Optional extension:** bound the pending queue as well as the active workers.

### 27. Cancellable batch CLI

**Time:** 4–6 hours. **Requires:** 26. **Practise:** context, shutdown, error propagation.

**Build:** let a user cancel a running batch.

1. Create an application-lifetime context that receives Ctrl-C or an explicit cancel command.
2. Stop submitting new files when it is cancelled.
3. Make workers check cancellation between units of supported work and return a clear cancelled result.
4. Drain or otherwise account for already-started jobs so no goroutine blocks on result delivery.
5. Test cancellation with channel coordination and test normal completion separately.

**Done when:** cancellation returns promptly, reports which jobs completed, and leaves no test worker blocked.

**Optional extension:** allow a per-file processing timeout. Explain that checking a context between CSV rows cannot interrupt every kind of blocked file read.

**Weekly checkpoint:** diagram or explain producer, workers and collector in a paragraph. State exactly who owns each channel and what stops the program. Run the race detector again on the tests that execute this path.

## Week 8 — Move work out of the request

**Read first:** state transitions, SQL transactions and request context lifetimes. Keep one application process for DataWatch.

### 28. Job-state simulator

**Time:** 3–4 hours. **Requires:** 08. **Practise:** state modelling and invariants.

**Build:** a tiny in-memory scheduler for queued, running, succeeded, failed and interrupted jobs.

1. Define allowed transitions in a small table.
2. Write one function that attempts a transition and rejects illegal changes.
3. Store created and last-updated times.
4. Simulate a queue, a worker finishing and a worker failing.
5. Test duplicate completion, failure after success and an unknown job ID.

**Done when:** a succeeded job cannot later become running through the normal transition function.

**Optional extension:** record a short failure reason.

### 29. Durable scan queue

**Time:** 5–7 hours. **Requires:** 19 and 26–28. **Practise:** reliable acknowledgement, bounded work, SQL state.

**Build:** DataWatch accepts an upload quickly and processes it in a bounded background pool.

1. Store the input under a generated ID, create a queued run in PostgreSQL and acknowledge only after both are saved.
2. Have a worker claim queued runs, transition them to running and process each file.
3. Write the final report and succeeded status as one consistent database operation; store a controlled failure reason on error.
4. Expose a status endpoint that returns the run ID and current state.
5. Test that two uploads can be queued while a worker is busy and that the queue limit produces a predictable response.

**Done when:** a slow scan leaves the HTTP API responsive, and the accepted queued run can be seen after restarting the application.

**Optional extension:** add an application-level timeout for each job. Workers need an application-lifetime context, not the upload request context, because the latter ends with the response.

### 30. Restart and retry drill

**Time:** 5–7 hours. **Requires:** 29. **Practise:** recovery and idempotence.

**Build:** make DataWatch's interrupted runs visible and safe to retry.

1. Decide how a run left in running state is marked interrupted at startup.
2. Add an explicit retry operation that transitions an interrupted run back to queued.
3. Ensure retries replace or uniquely identify one final report, so repeating a request cannot create duplicate committed reports.
4. Stop the process while a test job is active, restart it and inspect the persisted status.
5. Retry once and verify that the final result and history are coherent.

**Done when:** an interrupted job is never quietly shown as complete, and retrying it produces a single final report.

**Optional extension:** reconcile abandoned temporary files after a failed upload.

**Weekly checkpoint:** explain why an accepted upload must be persisted before the success response. Demo a slow job, a restart and an explicit retry using the same small input fixture.

## Week 9 — Let someone else use it

**Read first:** html/template, accessible form labels and HTTP form parsing. Start with clear tables and text; add visual polish after the facts are right. [R14]

### 31. HTML report exporter

**Time:** 3–4 hours. **Requires:** 12. **Practise:** templates and presentation of structured results.

**Build:** turn one JSON profile into a standalone HTML report.

1. Create a small template showing dataset name, record count and column summaries.
2. Read a profile into a Go struct and render the template.
3. Make missing, zero and unavailable numeric values visibly different.
4. Include an empty-state message for a header-only dataset.
5. Open the output in a browser and check it with data containing characters such as < and &.

**Done when:** the report remains readable for an empty file and a file with missing values, and dataset text does not become executable HTML.

**Optional extension:** add a small stylesheet to the exported document.

### 32. Upload dashboard

**Time:** 5–7 hours. **Requires:** 16, 20 and 29. **Practise:** browser forms, multipart uploads, state display.

**Build:** a simple DataWatch web view for uploading and browsing reports.

1. List datasets and their latest run status on a landing page.
2. Add a labelled upload form that accepts a CSV file and selects or creates a dataset.
3. Handle multipart uploads under a total request limit and store files through the existing upload service.
4. Redirect to the run detail page; display queued, running, failed and succeeded states clearly.
5. Add a report table that shows columns, missing counts and numeric summaries.

**Done when:** a person can upload the clean fixture in a browser and find its report without using curl or SQL.

**Optional extension:** use minimal JavaScript to refresh a pending run's status. Keep the basic page useful without it.

### 33. Comparison and rule view

**Time:** 4–6 hours. **Requires:** 22–23 and 32. **Practise:** communicating findings.

**Build:** show a run's rule violations and compare it with an explicitly selected baseline.

1. Add a small rule form to the upload flow for required columns, numeric ranges and missing-value limits; validate it before starting a run.
2. Display the saved rule configuration and each violation beside the finished report.
3. Provide a baseline selector containing previous runs of the same dataset.
4. Render added/removed columns and differences in counts and missing fractions; label unavailable values and percentage-point changes.
5. Ask another person to attempt the clean-then-corrupt demonstration; fix the point where they hesitate.

**Done when:** someone can define a rule, upload a changed file, see which rule failed and why, then compare it with their chosen baseline.

**Optional extension:** add a simple chart of missing fractions with text labels and a table fallback.

**Weekly checkpoint:** complete the demonstration through the browser: clean CSV, corrupted CSV, rule findings and selected-baseline comparison. Note the first usability issue a fresh user encounters.

## Week 10 — Make failures understandable

**Read first:** environment-based configuration, structured logging and graceful HTTP shutdown. Exercise a full application flow rather than only individual functions.

### 34. Config and health service

**Time:** 3–4 hours. **Requires:** 19. **Practise:** environment parsing, startup validation, operational status.

**Build:** make application settings and operational checks predictable.

1. Define settings for listen address, database connection, upload limit, data directory and worker count.
2. Read them from environment variables or a local config file, with documented defaults where appropriate.
3. Reject invalid settings before accepting requests.
4. Provide liveness and readiness routes; the readiness route checks dependencies needed to serve real work.
5. Test a missing database and an invalid worker count.

**Done when:** a fresh clone has a clear example configuration, and a database outage makes readiness fail without making the process look as though a real scan succeeded.

**Optional extension:** print a safe startup summary that omits credentials.

### 35. Logged and stoppable server

**Time:** 4–6 hours. **Requires:** 27, 29 and 34. **Practise:** structured logs, signals, graceful shutdown.

**Build:** make a DataWatch run easy to diagnose and stop.

1. Log upload accepted, job started, job finished or failed, with run ID and duration.
2. Avoid logging whole CSV bodies or secrets.
3. Stop accepting new requests on a termination signal.
4. Let active work finish within a chosen shutdown period or mark it for recovery.
5. Test an active job and a stop signal, then inspect the final persisted status.

**Done when:** a run can be traced through logs by its ID and stopping the server does not leave a completed run labelled running.

**Optional extension:** include a request ID in API errors and logs.

### 36. Reproducible setup and failure drill

**Time:** 6–8 hours. **Requires:** 30 and 32–35. **Practise:** integration, containers, recovery.

**Build:** set up the application and PostgreSQL with Compose and document one complete smoke test.

1. Package the app and database with persistent data volumes, migrations and a safe example environment file.
2. Start from a clean checkout and run the clean/corrupted CSV workflow.
3. Verify reports survive a restart of the app and database containers.
4. Exercise malformed CSV, oversized upload, full queue, database outage and interruption during a job.
5. For each failure, check the HTTP result, job state and log message; fix any inconsistent outcome.

**Done when:** a second person can follow the README to run DataWatch locally and observe one controlled example of a failed scan.

**Optional extension:** run database integration tests in CI with an isolated database.

**Weekly checkpoint:** record a real bug found by one of the drills, explain why it occurred and link the fix. Check the README against an actual clean checkout.

## Week 11 — Measure instead of guessing

**Read first:** Go's benchmark, fuzzing and diagnostics guides. Record the workload and conditions before changing performance-sensitive code. [R12, R15]

### 37. Benchmark fixture laboratory

**Time:** 5–7 hours. **Requires:** 12 and 26. **Practise:** benchmark design, generated test data, measurement.

**Build:** generate repeatable CSV fixtures and measure profile processing.

1. Produce at least small and medium inputs with specified row count, width and missing-value pattern; record the generator settings and fixed seed.
2. Verify the generated files have expected headers and record counts.
3. Benchmark the profiler alone using go test -bench, then record allocations with -benchmem.
4. Compare a one-worker and multi-worker batch with the same input, using repeated runs on the same machine.
5. Keep input files and the benchmark method stable while changing one implementation detail.

**Done when:** a reader can reproduce your measured comparison and tell whether the timing includes disk, database or HTTP work.

**Optional extension:** collect a CPU or heap profile and explain one expensive code path. Report an unfavourable result honestly if more workers make it slower.

### 38. Fuzz and regression lab

**Time:** 4–6 hours. **Requires:** 12 and 37. **Practise:** unexpected inputs, regression cases, profiling.

**Build:** exercise CSV profiling with generated inputs and fix one real issue it exposes.

1. Write deterministic tests for the tricky known cases first.
2. Add a fuzz target whose invariant is that any bytes produce either a controlled error or a valid, internally consistent result and never a panic.
3. Seed it with clean CSV, quoted multiline fields and malformed rows.
4. Run fuzzing for a bounded session; keep any interesting failure as a regression case and fix the cause.
5. Use a profile or allocation measurement to revisit one genuine bottleneck; record before and after results and unchanged output.

**Done when:** the repository includes a reproducible fuzz target, at least one meaningful regression case if a bug was found, and measurements for the performance change you chose.

**Optional extension:** fuzz the rule-file validator with a consistency invariant.

**Weekly checkpoint:** document Go version, input size, commands, repeated timings and limitations. Explain one correctness bug and one performance result with evidence.

## Week 12 — Finish a portfolio-quality demonstration

**Read first:** your own README as though you had never seen the repository. Choose a clean and corrupted sample with known, documented results.

### 39. End-to-end demo runner

**Time:** 4–6 hours. **Requires:** 30, 33 and 36. **Practise:** scripting, client polling, final verification.

**Build:** a small Go program that demonstrates DataWatch from the public API.

1. Accept a server address and the clean/corrupted sample paths as arguments.
2. Upload the clean file, record the returned run ID and poll until it reaches a terminal state.
3. Upload the corrupted file with the intended rule configuration and choose the clean run as baseline.
4. Print the rule failures and comparison summary; return a failing exit code if the API or expected outcome is wrong.
5. Run it against a fresh local deployment and document the exact command.

**Done when:** a reviewer can reproduce the entire story with one documented command after starting the service.

**Optional extension:** let the runner print a short machine-readable summary for CI.

### 40. DataWatch terminal client

**Time:** 4–6 hours. **Requires:** 14, 29 and 39. **Practise:** designing a usable command-line client.

**Build:** a separate Go executable that talks to the finished DataWatch API.

1. Add subcommands for upload, status, report and comparison.
2. Parse flags and validate missing arguments before making a request.
3. Use an http.Client with timeouts and show controlled responses for network, HTTP and malformed-JSON failures.
4. Display progress while a run is pending, then show its final report and quality status.
5. Test the client against httptest and try it manually against the local DataWatch server.

**Done when:** someone can inspect an upload from a terminal using only the documented CLI and API, without opening a browser or querying the database.

**Optional extension:** add JSON output suitable for piping into another command.

**Release checkpoint:** publish a clear README, screenshots or a short recording, sample data origins, local setup, test commands, measured benchmark results and honest limitations. Tag the application release once the documented workflow works from a fresh clone. If hosting a demo, use approved sample data and verify persistent state, HTTPS and the upload policy. Practise explaining a request flow, a restart failure, a database choice, a test that caught a bug and a benchmark result in your own words.

## 10 optional projects

These are independent choices for days when you want a change of subject. They do not block the 40-project route.

### B01. Hangman (after week 2)

**Build steps:** (1) Choose a word from a small built-in list and show hidden letters. (2) Accept one letter at a time, showing correct positions and remaining guesses. (3) Handle repeated guesses and define whether they cost an attempt.

**Done when:** the player can win or lose and repeated letters follow your rule. **Practise:** strings, runes, slices, state.

### B02. Maze solver (after week 2)

**Build steps:** (1) Represent a small grid with walls, start and exit. (2) Find a path using breadth-first search and a queue. (3) Print the path, and return “no route” for a blocked map.

**Done when:** the shortest route in a hand-drawn tiny maze matches the output. **Practise:** queues, maps, indexing.

### B03. Job-application tracker (after week 3)

**Build steps:** (1) Store company, role, application date and status in JSON. (2) Add commands to record an application and change its status. (3) List entries by status and next follow-up date.

**Done when:** entries survive restarts and invalid status changes are rejected. **Practise:** persistence, state transitions, dates.

### B04. Football league table (after week 3)

**Build steps:** (1) Parse a fixture CSV with home team, away team and scores. (2) Award 3 points for a win, 1 for a draw and 0 for a loss. (3) Sort teams by points, goal difference, goals scored and then name as an explicit exercise rule.

**Done when:** a five-match fixture with hand-calculated standings produces the expected table. **Practise:** structs, aggregation, sorting.

### B05. CBZ comic catalogue (after week 3)

**Build steps:** (1) Open a small .cbz ZIP archive and list its files. (2) Count recognised page-image entries and sort their names. (3) Summarise title, page count and obvious gaps in page numbers without extracting images.

**Done when:** an archive with three pages and an unrelated text file reports three pages. **Practise:** archive/zip, file metadata, filtering.

### B06. Markdown link checker (after week 4)

**Build steps:** (1) Begin with a list of URLs extracted manually from a Markdown file. (2) Request each address with a timeout and record status or error. (3) Add simple Markdown link extraction for a defined subset of syntax and report broken links with line numbers.

**Done when:** a local httptest page with a 404 is reported alongside a successful page. **Practise:** text parsing, HTTP clients, test servers.

### B07. RSS reader (after week 4)

**Build steps:** (1) Parse a local XML feed fixture into a small struct. (2) Print item title, link and date in readable order. (3) Fetch a configured feed with an HTTP timeout and report malformed XML.

**Done when:** the local fixture works without internet access and a bad feed produces an error. **Practise:** encoding/xml, HTTP, sorting.

### B08. URL shortener (after week 5)

**Build steps:** (1) Build a local POST route that stores a URL and generated short code in PostgreSQL. (2) Redirect from a GET route using the code. (3) Reject an invalid URL and show a clear result for an unknown code.

**Done when:** a stored link still redirects after restarting the program. **Practise:** handlers, SQL, response headers.

### B09. Backup integrity verifier (after week 3)

**Build steps:** (1) Write a manifest of relative file paths, sizes and SHA-256 hashes for a chosen directory. (2) Recompute them on a later run. (3) Report added, changed, removed and unreadable files separately.

**Done when:** changing one byte in a file gives a “changed” result while untouched files stay verified. **Practise:** file walking, hashing, comparison.

### B10. Concurrent cache (after week 7)

**Build steps:** (1) Implement a map of keys to values with a set/get interface. (2) Add expiry times and ensure an expired value is not returned. (3) Use a mutex around shared state and test concurrent readers and writers under go test -race.

**Done when:** an expired entry is absent and concurrent tests exercise the cache without a data race. **Practise:** mutexes, time, race detection.

## Weekly reflection template

Keep this short in your notes:

- **Built:** link to finished program or DataWatch commit.
- **Can explain:** one Go concept in your own words.
- **Bug:** the input that exposed it, root cause and fix.
- **Next:** one concrete action to start the next session.
- **Ready to move on?** demonstrate the week's checkpoint, then choose yes or repeat a difficult card.

At the end of weeks 4, 6, 8 and 12, record a 2–5 minute screen demo. An unfinished but honest progress record is useful both for learning and for explaining the project during applications.

## References

The project ideas, sequence and completion checks here are curriculum choices. Read the official material as you encounter each feature; the links were checked when the curriculum was written.

- [R1] A Tour of Go: https://go.dev/tour/welcome/1
- [R2] How to Write Go Code: https://go.dev/doc/code
- [R3] Testing package: https://pkg.go.dev/testing
- [R4] io package: https://pkg.go.dev/io
- [R5] CSV package: https://pkg.go.dev/encoding/csv
- [R6] HTTP package: https://pkg.go.dev/net/http
- [R7] HTTP testing package: https://pkg.go.dev/net/http/httptest
- [R8] Go database guide: https://go.dev/doc/database/
- [R9] PostgreSQL tutorial: https://www.postgresql.org/docs/current/tutorial.html
- [R10] Context package: https://pkg.go.dev/context
- [R11] Go race detector: https://go.dev/doc/articles/race_detector
- [R12] Go fuzzing tutorial: https://go.dev/doc/tutorial/fuzz
- [R13] math/rand/v2: https://pkg.go.dev/math/rand/v2
- [R14] HTML template package: https://pkg.go.dev/html/template
- [R15] Diagnostics and profiling: https://go.dev/doc/diagnostics

**Begin with project 01:** write three example unit conversions, implement them as functions, then make a small CLI. Move to project 02 once an invalid input produces a useful error you understand.
