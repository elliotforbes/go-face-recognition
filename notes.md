Face Recognition
==================

In this tutorial, we'll be implementing face recognition and detection using local binary patterns.

## Face Detection and Face Recognition

The process of recognizing a person from an image tends to follow a 3 step process. First step is face representation, the second is feature extraction and the final step is classification.

* Face Representation - This is how you model a face.
* Feature Extraction - In this phase we extract the most useful and unique features (properties) of a face image
* Classification - With the most useful properties of the face extracted, we move into the classification phase and we try to compare these properties against an image database.

## LBPH

LBP is a type of visual descriptor used for classification in computer vision. It typically subdivides the face area of an image into small regions from which Local Binary Patterns (LBP), histograms are extracted and concatenated into a single feature vector.

#### Histograms

Histograms are accurate representations of the distribution of numerical data. It is an estimate of the probability distribution of a continuous variable(quantitative variable). If you had a range of values for one particular variable, you would typically start by bin-ing or bucket-ing the range of values. This divides the entire range of values into a series of intervals, you then proceed to count how many values fall into each interval.

Say, for instance, I had 500 values that ranged from -3.5 to 3.5. We could `bin` these items like so:

| bin | Count |
|--|--|
| -3.5 to -2.51 | 9 |
| -2.5 to -1.51 | 32 |
| -1.5 to -0.51 | 109 |
| -0.5 to 0.49 | 180 |
| 0.5 to 1.49 | 132 |
| 1.5 to 2.49 | 34 |
| 2.5 to 3.49 | 4 |

This would then produce the following histogram for these values:

![](https://upload.wikimedia.org/wikipedia/commons/thumb/1/1d/Example_histogram.png/220px-Example_histogram.png)


## References

1. Face Recognition using Local Binary Patterns (LBP) - https://globaljournals.org/GJCST_Volume13/1-Face-Recognition-using-Local.pdf
