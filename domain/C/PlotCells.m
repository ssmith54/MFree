close all
clear all 


addpath('Cells')
set(0,'DefaultTextFontname', 'latex')
set(0,'DefaultAxesFontName', 'latex')


figure
filename = "cells.txt";
subplot(1,2,1)
[lShape, ~, ~ ] = read_polygon(filename);
fill(lShape(:,1),lShape(:,2),'y');
xlim([min(lShape(:,1)) - 0.5, max(lShape(:,1)) + 0.5])
ylim([min(lShape(:,2)) - 0.5, max(lShape(:,2)) + 0.5])
axis equal
area = polyarea(lShape(:,1),lShape(:,2))

subplot(1,2,2)



hold on


files = dir('Cells/*.txt');
for file = files'
    [poly, ~, ~ ] = read_polygon(file.name);
    hold on
    fill(poly(:,1),poly(:,2),rand(1,3));
    % Do some stuff
end


ylim([0,1])
axis equal



