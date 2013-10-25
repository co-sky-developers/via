via - fast shortest path computation
====================================

**via** is a new, lightweight shortest path problem computation service fully implemented using RESTful techniques. via provides distance matrix calculation through a simple API.

Installing via
--------------

Simply clone this repository and run ``go build -ldflags="-r ."``. The current working directory is required to be added in ``LD_LIBRARY_PATH`` because of the requirements of the generated
C++ code from contraction hierarchies. See requirements for what you need to install.

via includes source code of the contraction hierarchies project, see [here](http://algo2.iti.kit.edu/routeplanning.php) for more information, which is also licensed under the AGPL.

Requirements
------------

* Go 1.2 
* Redis for fast key-value storage
* Swig 2.0.11 for C++ wrapping
* PostgreSQL 
* Preprocessed OpenStreetMap data using contraction hierarchies. We do not currently provide these, but instructions on how to compile them yourself will be given in the future.

About
-----

via is a project by the CO-SKY research team at the University of Jyväskylä. 

http://www.co-sky.fi

License
-------

via is redistributed using the GNU Affero General Public License, version 3 (19th Nov 2007). See LICENSE for further instructions for redistribution.
