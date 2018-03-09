# Rest api for hn.algolia.com queries

This is a Rest api to provides data about the reqest made on the hn.algolia.com webservice.

The data is expected to be in a tsv file in the folder where the application is started.
I did not include it into the repository as it doesn't belong to me.

# Implementation
The challenging part of the application is to process and index in an efficient way the somewhat big tsv file (1623420 lines).

After some consideration I decided to use simple slices as the data structure to store the values from the tsv.
The timestamp column and the query column are each stored in it's own slice which are then sorted based on the timestamp.
The timestamp are stored as string, this way it's easy to get range by getting the first and last string that have a given query (like `2015-03`) as a substring.
An advantage of that is that I can match partial timerange `2015-03-1` to get the days between the 10 and the 20 of Mars.
This datastructure doesn't allow direct retrieval of one point of data but values corresponding to a given timerange are stored continuously.
So event if a binary search is required to get the starting and end index of a timerange the matching queries can then be passed directly as a slice.

To get the distinct values from a list of queries I use a set and iterate on the list.
This is a O(n) operation on the list.

To count the popularity I first use a dictionnary to count the frequencies of each query.
Then I implemented quickselect as a way to get the K biggest values without sorting all the queries.

I think this implementation works well for the small TSV I have corresponding only to a few days of logs.
It is an easy index system which I've implemented to get a feeling of the dataset and kept since it worked well for the size of the data I have.

However if the service was to be used with a lot more data it might be necessary to improve the data structure to serve more rapidly.
One solution would be to build an index giving for example the coordinate corresponding to the range of each day in the index.
Then query on a smaller time interval wouldn't have to query the whole index.

The other part that could quickly get out of hand is queries requesting popularity values on a time period too big.
The solution would probably be to store those values separetly.
So maybe build a list of the most popular queries for each day, those list could be merged for queries regarding a bigger timerange.

Another option is to divide the workload, some parrallel processing is possible, for exemple when looking for distinct queries it's possible to get the distinct values on sublist in parrallel and then merge those.
