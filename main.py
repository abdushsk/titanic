import pandas as pd
from sklearn.tree import DecisionTreeClassifier  # Import Decision Tree Classifier
# Import train_test_split function
from sklearn.model_selection import train_test_split
# Import scikit-learn metrics module for accuracy calculation
from sklearn import metrics
import numpy as np

df = pd.read_csv("assets/surv.csv")

print(df.dropna(inplace=True))

df["Sex"].replace("male", 1, inplace=True)
df["Sex"].replace("female", 2, inplace=True)

df["Embarked"].replace("S", 1, inplace=True)
df["Embarked"].replace("C", 2, inplace=True)
df["Embarked"].replace("Q", 3, inplace=True)


print(df.isna().sum())


X = df[["Pclass", "Sex", "Age", "SibSp", "Parch", "Fare", "Embarked"]]

y = df.Survived

# print(X)

X_train, X_test, y_train, y_test = train_test_split(
    X, y, test_size=0.5, random_state=1)

clf = DecisionTreeClassifier(
    max_depth=5, criterion="entropy", min_samples_split=50)

clf.fit(X_train, y_train)

# Compute the cost-complexity pruning path
path = clf.cost_complexity_pruning_path(X_train, y_train)
ccp_alphas, impurities = path.ccp_alphas, path.impurities

# Find the optimal value for ccp_alpha
clfs = []

for ccp_alpha in ccp_alphas:
    print(ccp_alpha)
    clf = DecisionTreeClassifier(
        random_state=0, max_depth=5, criterion="entropy", min_samples_split=50, ccp_alpha=ccp_alpha)
    clf.fit(X_train, y_train)
    clfs.append(clf)

# Predict the response for test dataset
highest = 0.0

acc_scores = [metrics.accuracy_score(
    y_test, clfa.predict(X_test)) for clfa in clfs]

print(max(acc_scores))
