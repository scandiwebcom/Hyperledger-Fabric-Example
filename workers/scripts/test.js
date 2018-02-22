'use strict';
/*
* Copyright IBM Corp All Rights Reserved
*
* SPDX-License-Identifier: Apache-2.0
*/
/*
 * Chaincode query
 */
if (process.argv.length <= 2) {
    console.log("Usage: " + __filename + " SOME_PARAM");
    process.exit(-1);
}

const config = require('./config.js');
const IP = config.ip;

/** 
 * Get arguments
 * 0 - Command
 * 1 - Filename
 * 2 - Execution method
 * 3 - Data. e.g. to create new worker "WORKER12, Joe, working, QA"
 */
const args = process.argv.slice(2);
const method = args[0];
const data = args[1] ? args[1].split(",") : [];

/**
 * query          - node test.js query
 * changePosition - node test.js changePosition "WORKER0,TEAM LEAD"
 * queryWorker    - node test.js queryWorker "WORKER0"
 * addWorker      - node test.js addWorker "WORKER22,Joe,working,Developer"
 */
if (method === 'query') {
    const run = require('./query.js')(config, data);
    run();
} else if (method === 'changePosition') {
    const run = require('./changePosition.js')(config, data);
    run();
} else if (method === 'queryWorker') {
    const run = require('./queryWorker.js')(config, data);
    run();
} else if (method === 'addWorker') {
    const run = require('./addWorker.js')(config, data);
    run();
} else {
    console.info('Execution file not found.');
}

console.info();