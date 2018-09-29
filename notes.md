Face Recognition
==================

In this tutorial, we'll be implementing face recognition and detection using local binary patterns.

## The Process

The process of recognizing a person from an image tends to follow a 3 step process. First step is face representation, the second is feature extraction and the final step is classification.

* Face Representation - This is how you model a face.
* Feature Extraction - In this phase we extract the most useful and unique features (properties) of a face image
* Classification - With the most useful properties of the face extracted, we move into the classification phase and we try to compare these properties against an image database.

## Feature Extraction with Local Binary Patterns

In order to extract some of the most useful features from an image of a face, we can use the Local Binary Pattern method.  

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

## LBP

The Local Binary Pattern looks at 9 pixels at a time. It then takes these 9 pixel values and computes a single value from it. It takes the central value and it compares it against all of it's neighbouring values and if the neighbouring value is higher, it replaces it with a 1, if it is lower, it is replaced with a 0. These 8 neighbour bits are then turned into a byte.

| 1 | 1 | 0 |
| - | - | - |
| 0 | - | 1 |
| 0 | 1 | 0 |

Equates to [1,1,0,1,0,1,0,0] if we turned this into an array of Bytes.

## References

1. Face Recognition using Local Binary Patterns (LBP) - https://globaljournals.org/GJCST_Volume13/1-Face-Recognition-using-Local.pdf
